package com.example.whq_android;

import androidx.appcompat.app.AppCompatActivity;

import android.content.Intent;
import android.os.Bundle;
import android.view.View;
import android.widget.Button;
import android.widget.EditText;
import android.widget.TextView;

public class ChatActivity extends AppCompatActivity {

    private Button btn_send;
    private EditText et_Msg;
    private TextView tv_Msg;

    private Listen listen;
    private Client client;
    private boolean isForm = true;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_chat);
        Intent intent = getIntent();
        client = (Client) intent.getSerializableExtra("Conn");
        init();
        btn_send.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                client.SendMessage(et_Msg.getText().toString());
            }
        });
    }

    private void init(){
        btn_send = findViewById(R.id.btn_send);
        et_Msg = findViewById(R.id.SendMessage);
        tv_Msg = findViewById(R.id.tv_msg);
    }

    class Listen extends Thread{
        @Override
        public void run() {
            super.run();
            while(isForm){
                String accept = client.ListenMessage();
                runOnUiThread(new Runnable() {
                    @Override
                    public void run() {
                        tv_Msg.setText(tv_Msg.getText().toString()+accept);
                    }
                });
            }
        }
    }

    protected void onResume() {
        super.onResume();
        listen = new Listen();
        isForm = true;
        listen.start();
    }

    protected void onPause() {
        super.onPause();
        isForm = false;
    }

}