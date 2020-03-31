# win10 python2.75 安装pip报错

1. 错误内容：`after connection broken by 'SSLError(SSLError(1, '_ssl.c:504: error:1407742E:SSL routines:SSL23_GET_SERVER_HELLO:tlsv1 alert protocol version'),)': /simple/pip/`

2. 解决办法：`python get-pip.py -i http://pypi.douban.com/simple/ --trusted-host pypi.douban.com`

   1. 详见：<https://blog.csdn.net/u012425536/article/details/89304645>

3. 类似的：

   ```shell
   python.exe -m pip install -U "pylint<2.0.0" --user  -i http://pypi.douban.com/simple/ --trusted-host pypi.douban.com
   python.exe -m pip install -U "numpy<=1.16.1" --user  -i http://pypi.douban.com/simple/ --trusted-host pypi.douban.com
   python.exe -m pip install -U "gensim" --user  -i http://pypi.douban.com/simple/ --trusted-host pypi.douban.com
   
   # 安装mysql三步
   python.exe -m pip install -U "wheel" --user  -i http://pypi.douban.com/simple/ --trusted-host pypi.douban.com
   # https://download.lfd.uci.edu/pythonlibs/n5jyqt7p/backports.lzma-0.0.13-cp27-cp27m-win_amd64.whl 下载所需文件
   python.exe -m pip install -U C:\Users\admin\Downloads\mysqlclient-1.4.4-cp27-cp27m-win_amd64.whl
   
   MySQLdb
   ```

   

