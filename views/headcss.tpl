{{define "headcss"}}

    *{
      margin: 0px;
      padding: 0px;
    }
    
    #top{width:100%;height: 50px;background: rgba(50,50,50,0.5);position: fixed;float: left; top:0px;}
    #top .top_first{float: left; color: red;height: 50px;width: 150px;line-height: 50px;text-align: center;  cursor: pointer;position: relative;}
    #top .top_first:hover{background:#90c;}
    #top .top_first .top_first_list{width: 150px;height: 250px;position: absolute;top: 50px;background: blue; background:rgba(0,0,0,0.9);display: none;}
    #top .top_first .top_first_list ul li{width: 150px; color: yellow;}
    #top .top_first .top_first_list ul li a{text-decoration:none;color: yellow;width: 150px;height: 50px;display: block;}
    #top ul li{float: left; color: red;height: 50px;width: 100px;display: block;line-height: 50px;text-align: center;  cursor: pointer;}
    #top ul li:hover{background:#90c;}
    #top .logo{height: 45px;width: 80px; background-image: url("../static/img/logo.png"); background-size: cover;margin: 20px;margin-top:2px;float: left; }
    .menu{background:#90c;}
{{end}}