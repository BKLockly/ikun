# ikun
> 超级简单的鸡你太美cmd字符画，开箱即用

![image-20230302160401989](https://img2023.cnblogs.com/blog/3038812/202303/3038812-20230302160403315-98896569.png)



### 说明

很简单的实现，音乐同步和帧率都没有做，就纯纯图一乐。

### 使用

1. git clone https://github.com/BKLockly/ikun.git
2. 直接运行exe
3. 调整cmd界面，字体改为6，宽度与高度如下：

![image-20230302160524058](https://img2023.cnblogs.com/blog/3038812/202303/3038812-20230302160523998-519327298.png)



### 其他视频的制作

1. 将准备好的视频文件逐帧导出，有很多方式，这里直接推荐最简单的软件：[Free Video to JPG Converter](https://www.dvdvideosoft.com/products/dvd/Free-Video-to-JPG-Converter.htm)

   ![image-20230302161536271](https://img2023.cnblogs.com/blog/3038812/202303/3038812-20230302161723119-1013993325.png)

2. 转ascll用的这个 [Ascii Generator 2](https://sourceforge.net/projects/ascgen2/)

   导入转化过的文件夹，长宽高调整一下，后续终端界面的设置与这里相同

   ![image-20230302161941110](https://img2023.cnblogs.com/blog/3038812/202303/3038812-20230302161941120-1577208043.png)

3. 将go文件放入导出txt的目录下面，修改里面的`pageLimit`为图片数，还有长宽高的常量。目前写的很垃圾，就用的`time.Sleep(22 * time.Millisecond)`，主要还是太菜了本地配gocv老是不成功，后续有时间再更这个stupid code

   

   



