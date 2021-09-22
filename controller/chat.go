package ctrl

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
)

const (
	CMD_SINGLE_MSG = 10
	CMD_ROOM_MSG   = 11
	CMD_HEART      = 0
)

type Message struct {
	Id      int64  `json:"id,omitempty" form:"id"`           //消息ID
	Userid  int64  `json:"userid,omitempty" form:"userid"`   //谁发的
	Cmd     int    `json:"cmd,omitempty" form:"cmd"`         //群聊还是私聊
	Dstid   int64  `json:"dstid,omitempty" form:"dstid"`     //对端用户ID/群ID
	Media   int    `json:"media,omitempty" form:"media"`     //消息按照什么样式展示
	Content string `json:"content,omitempty" form:"content"` //消息的内容
	Pic     string `json:"pic,omitempty" form:"pic"`         //预览图片
	Url     string `json:"url,omitempty" form:"url"`         //服务的URL
	Memo    string `json:"memo,omitempty" form:"memo"`       //简单描述
	Amount  int    `json:"amount,omitempty" form:"amount"`   //其他和数字相关的
}

/**
消息发送结构体
1、MEDIA_TYPE_TEXT
{id:1,userid:2,dstid:3,cmd:10,media:1,content:"hello"}
2、MEDIA_TYPE_News
{id:1,userid:2,dstid:3,cmd:10,media:2,content:"标题",pic:"http://www.baidu.com/a/log,jpg",url:"http://www.a,com/dsturl","memo":"这是描述"}
3、MEDIA_TYPE_VOICE，amount单位秒
{id:1,userid:2,dstid:3,cmd:10,media:3,url:"http://www.a,com/dsturl.mp3",anount:40}
4、MEDIA_TYPE_IMG
{id:1,userid:2,dstid:3,cmd:10,media:4,url:"http://www.baidu.com/a/log,jpg"}
5、MEDIA_TYPE_REDPACKAGR //红包amount 单位分
{id:1,userid:2,dstid:3,cmd:10,media:5,url:"http://www.baidu.com/a/b/c/redpackageaddress?id=100000","amount":300,"memo":"恭喜发财"}
6、MEDIA_TYPE_EMOJ 6
{id:1,userid:2,dstid:3,cmd:10,media:6,"content":"cry"}
7、MEDIA_TYPE_Link 6
{id:1,userid:2,dstid:3,cmd:10,media:7,"url":"http://www.a,com/dsturl.html"}

7、MEDIA_TYPE_Link 6
{id:1,userid:2,dstid:3,cmd:10,media:7,"url":"http://www.a,com/dsturl.html"}

8、MEDIA_TYPE_VIDEO 8
{id:1,userid:2,dstid:3,cmd:10,media:8,pic:"http://www.baidu.com/a/log,jpg",url:"http://www.a,com/a.mp4"}

9、MEDIA_TYPE_CONTACT 9
{id:1,userid:2,dstid:3,cmd:10,media:9,"content":"10086","pic":"http://www.baidu.com/a/avatar,jpg","memo":"胡大力"}

*/

type Node struct {
	Conn *websocket.Conn
	// 并行转串行
	DataQueue chan []byte
	GroupSets set.Interface
}

var clientMap map[int64]*Node = make(map[int64]*Node)
var rwlocker sync.RWMutex

func Chat(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")
	token := query.Get("token")
	userId, _ := strconv.ParseInt(id, 10, 64)
	isvalid := checkToken(userId, token)
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return isvalid
		},
	}).Upgrade(w, r, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
	// TODO: Get All userid in Group
	cmdIds := contactService.SearchCommunityIds(userId)
	for _, v := range cmdIds {
		node.GroupSets.Add(v)
	}
	rwlocker.Lock()
	clientMap[userId] = node
	rwlocker.Unlock()
	go sendproc(node)
	go recvproc(node)
	sendMsg(userId, []byte("hello there!~"))
}

func checkToken(userid int64, token string) bool {
	//从数据库里面查询并比对
	user := userService.Find(userid)
	return user.Token == token
}

func sendproc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}
	}
}

// {id,userid,dstid,cmd,media,content}
func recvproc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			return
		}
		//todo 对data进一步处理
		dispatch(data)
		// broadcast message to local
		broadMsg(data)
		fmt.Printf("Receiv <= %s \n", data)
	}
}

// listen to local port 3000 udp
func udprecvproc() {
	log.Println("Start udp receive process")
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3000,
	})
	defer conn.Close()
	if err != nil {
		log.Println(err.Error())
	}
	for {
		var buf [512]byte
		n, err := conn.Read(buf[0:])
		if err != nil {
			log.Println(err.Error())
			return
		}
		dispatch(buf[0:n])
	}
	log.Println("Stop udp receive process")
}

// send to local net other
func udpsendproc() {
	log.Println("Start udp send process")
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(192, 168, 3, 255),
		Port: 3000,
	})
	defer conn.Close()
	if err != nil {
		log.Println(err.Error())
		return
	}
	for {
		select {
		case data := <-udpsendchan:
			_, err := conn.Write(data)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}
	}
}

var udpsendchan chan []byte = make(chan []byte, 1024)

func broadMsg(data []byte) {
	udpsendchan <- data
}

func dispatch(data []byte) {
	//todo parse data into Message
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		log.Println(err.Error())
		return
	}
	//todo process Message according to cmd
	switch msg.Cmd {
	case CMD_SINGLE_MSG:
		sendMsg(msg.Dstid, data)
	case CMD_ROOM_MSG:
		//todo room char
		for _, v := range clientMap {
			if v.GroupSets.Has(msg.Dstid) {
				v.DataQueue <- data
			}
		}
	case CMD_HEART:
		//todo hearbeat,do nothing
	}
}

func sendMsg(userId int64, msg []byte) {
	//
	rwlocker.RLock()
	node, ok := clientMap[userId]
	rwlocker.RUnlock()
	if ok {
		node.DataQueue <- msg
	}
}
func AddGroupId(userId, groupId int64) {
	rwlocker.Lock()
	node, ok := clientMap[userId]
	if ok {
		node.GroupSets.Add(groupId)
	}
	rwlocker.Unlock()
}