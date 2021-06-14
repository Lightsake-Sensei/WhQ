package com.example.whq_android;

import androidx.appcompat.app.AppCompatActivity;

import android.content.Intent;
import android.os.Bundle;
import android.view.View;
import android.widget.Button;
import android.widget.EditText;
import android.widget.Toast;

import java.io.Serializable;

public class ConnectionActivity extends AppCompatActivity {

    private EditText Et_Ip,Et_Port;
    private Button btn_connect;

    private Check check;

    private boolean isForm = true;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_connection);
        init();
        check.start();

        btn_connect.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Toast.makeText(ConnectionActivity.this,"正在创建对"+Et_Ip.getText().toString()+":"+Et_Port.getText().toString()+"的连接",Toast.LENGTH_SHORT).show();
                int port = 1234;
                try{
                    port = Integer.parseInt(Et_Port.getText().toString());
                }catch (Exception e){
                    Toast.makeText(ConnectionActivity.this,"请输入正确格式的port",Toast.LENGTH_SHORT).show();
                }
                Client client = new Client(Et_Ip.getText().toString(),port);
                if(client.Connection()){
                    Toast.makeText(ConnectionActivity.this,Et_Ip.getText().toString()+":"+Et_Port.getText().toString()+"连接成功！",Toast.LENGTH_SHORT).show();
                }else{
                    Toast.makeText(ConnectionActivity.this,Et_Ip.getText().toString()+":"+Et_Port.getText().toString()+"连接失败,请检测设置是否正确",Toast.LENGTH_SHORT).show();
                    return;
                }
                Intent intent = new Intent(ConnectionActivity.this,ChatActivity.class);
                intent.putExtra("Conn", (Serializable) client);
                isForm = false;
                startActivity(intent);
            }
        });
    }

    //检查输入值
    class Check extends Thread{
        @Override
        public void run() {
            super.run();
            while(isForm){
                try {
                    sleep(500);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
                runOnUiThread(new Runnable() {
                    @Override
                    public void run() {
                        if(Et_Ip.getText().toString().length() > 0 && Et_Port.getText().toString().length() > 0 ){
                            btn_connect.setEnabled(true);
                        }else{
                            btn_connect.setEnabled(false);
                        }
                    }
                });
            }
        }
    }

    //初始化控件
    private void init(){
        Et_Ip = findViewById(R.id.Ip);
        Et_Port = findViewById(R.id.Port);
        btn_connect = findViewById(R.id.login);
        check = new Check();
    }
}