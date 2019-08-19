# windows远程桌面使用步骤

加上B要远程连接到A的电脑上，进行远程造作。

1. A输入`win+r`，调出运行命令的程序，输入`msra`，点击`邀请信任的人帮助你`，选择`将邀请另存为文件`，得到邀请文件和对应密码（一次性的）。
2. A将邀请文件和密码发送给B。
3. B点击A的邀请文件，会打开`mstsc`程序，输入对应密码
4. A接受B的远程连接请求，此时B可以看到A的电脑画面，但是不能操作。
5. B点击`mstsc`程序左上角的`请求控制`，A接受B的控制请求，B就可以操作A的电脑了。
6. 其他功能：

- B可以点击`mstsc`上排菜单的`实际大小/适应屏幕`按钮来切换A的屏幕画面在B的`mstsc`中显示模式（注意似乎只能缩小，不能放大；也即A的电脑分辨率小时，A的桌面画面不能占满B的屏幕；A的分辨率大时，A的桌面画面可以占满(适应屏幕)甚至超出B的屏幕(实际大小)）
- B可以点击`mstsc`上排菜单的`聊天`来与A进行聊天，A在A电脑上的`msra/Windows远程协助`程序上可以看到聊天信息，A也可以会信息。再次点击上排惨淡的`聊天`按钮可以关闭聊天界面。
