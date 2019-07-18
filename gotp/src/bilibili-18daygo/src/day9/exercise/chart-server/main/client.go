package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type Client struct {
	conn net.Conn
	buf  [8192]byte
}

func (p *Client) readPackage() (msg Message, err error) {
	n, err := p.conn.Read(p.buf[0:4])
	if n != 4 {
		err = errors.New("read header failed")
		return
	}

	buffer := bytes.NewBuffer(p.buf[0:4])
	var packLen uint32
	err = binary.Read(buffer, binary.BigEndian, &packLen)
	if err != nil {
		fmt.Println("read package len failed")
		return
	}

	n, err = p.conn.Read(p.buf[0:packLen])
	if n != int(packLen) {
		err = errors.New("read body failed")
		return
	}

	err = json.Unmarshal(p.buf[0:packLen], &msg)
	if err != nil {
		fmt.Println("unmarshal failed, err: ", err)
	}

	return
}

func (p *Client) Process() {
	for {
		msg, err := p.readPackage()
		if err != nil {
			return
		}

		err = p.processMsg(msg)
		if err != nil {
			return
		}
	}
}

func (p *Client) processMsg(msg Message) (err error) {
	switch (msg.Cmd) {
	case UserLogin:
		err = p.login(msg)
	case UserRegister:
		err = p.register(msg)
	default:
		err = errors.New("unsupport message")
		return
	}
	return
}

func (p *Client) login(msg Message) (err error) {
	defer func() {
		var respMsg Message
		respMsg.Cmd = UserLoginRes

		var loginRes LoginCmdRes
		if err != nil {
			loginRes.Code = 500
			loginRes.Error = fmt.Sprintf("%v", err)
		}

		data, err := json.Marshal(loginRes)
		if err != nil {
			fmt.Println("marchal failed")
			return
		}

	}()
	var cmd LoginCmd
	err = json.Unmarshal([]byte(msg.Data), &cmd)
	if err != nil {
		return
	}

	user, err := mgr.Login(cmd.Id, cmd.Passwd)
	if err != nil {
		return
	}
	return
}

func (p *Client) register(msg Message) (err error) {
	var cmd RegisterCmd
	err = json.Unmarshal([]byte(msg.Data), &cmd)
	if err != nil {
		return
	}

	user, err := mgr.Login(cmd.Id, cmd.Passwd)
	if err != nil {
		return
	}
	return
}

func (p *Client) writePackage(data []byte) (err error) {
	packLen := len(data)
	buffer := bytes.NewBuffer(p.buf[0:4])
	var packLen uint32
	err = binary.Write(buffer, binary.BigEndian, &packLen)
	if err != nil {
		fmt.Println("write package len failed")
		return
	}

	n, err := p.conn.Write(p.buf[0:4])
	if n != int(packLen) {
		err = errors.New("write data failed")
		return
	}

	n, err = p.conn.Write(data)
	if n != int(packLen) {
		err = errors.New("write data failed")
		return
	}
	if n != packLen {

	}

	return
}