TCP:面向连接，可靠的数据包传输

UDP:无连接的，不可靠的报文传递。网络状况不好，丢包严重。速度快。

UDP通信过程：
    
    1.创建用于通信的socket
    2.阻塞读socket
    3.处理读到的数据
    4.写数据给客户端

服务端主要用到函数

    1.创建监听地址：func ResolveUDPAddr(network,address string)
    
    2.创建用户通信的socket: func ListenUDP(network string, laddr *UDPAddr) (*UDPConn, error) {
    
    3.接收UDP数据：func ReadFromUDP(b []byte)
    
    4.写数据到UDP：func WriteToUDP(b []byte,addr *UDPAddr)

客户端用的函数：
    参考TCP

TCP：对不稳定的网络层，做完全弥补操作
UDP：-------------，不作为

使用场景：
    TCP：对数据传输安全性，稳定性较高的场合。网络文件的传输，下载，上传。
    UDP：对数据传输要求较高的场合。视频直播，再点电话会议。游戏。
