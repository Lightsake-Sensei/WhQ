package com.example.whq_android;

import android.widget.TextView;

import java.io.DataInputStream;
import java.io.DataOutputStream;
import java.io.IOException;
import java.net.Socket;

public class Client {
    private String Ip = "127.0.0.1";
    private int Port = 1234;
    private Socket conn;
    //空构造函数
    public Client(){

    }

    //带参构造函数
    public Client(String ip,int port){
        this.Ip = ip;
        this.Port = port;
    }

    //连接服务器
    public boolean Connection(){
        try {
            Socket s = new Socket(this.Ip, this.Port);
            this.conn = s;
        } catch (IOException e) {
            e.printStackTrace();
            return false;
        }
        return true;
    }

    public void SendMessage(String msg){
        // 向服务器端发送信息的DataOutputStream
        try {
            DataOutputStream out = new DataOutputStream(this.conn.getOutputStream());
            out.writeUTF(msg);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    public String ListenMessage(){
        try {
            // 读取服务器端传过来信息的DataInputStream
            DataInputStream in = new DataInputStream(this.conn.getInputStream());
            String accpet = in.readUTF();
            return accpet;
        } catch (IOException e) {
            e.printStackTrace();
        }
        return "";
    }

}
