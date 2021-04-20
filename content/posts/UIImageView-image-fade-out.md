---
title: UIImageView更换图片时渐隐渐现
date: 2015-11-02 11:47:44
tags: [iOS]
---

实现原理十分简单，使用UIImageView的透明度即可。然后在动画中完成。
alpha = 1 为全透明。
运行图例：

![UIImageView-image-fade-out](https://p3.ssl.qhimg.com/t0150836f31829840e4.gif)


实现代码:

    //图片渐隐渐现
    self.backgroundView.alpha = 0.7;
    [UIView animateWithDuration:0.5 animations:^{
        self.backgroundView.alpha = 1;
        self.backgroundView.image = [UIImage imageNamed:@"weather_bg_02.jpg"];

    }];
    

---

后来我知道了可以用更好的方法实现， 后续会有更新