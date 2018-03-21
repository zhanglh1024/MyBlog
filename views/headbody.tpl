{{define "headbody"}}
  <div id="top">
      <div class="logo"></div>
      <div class="top_first" ><span {{if .IsFrame}}class="menu"{{end}} style="display:block;">选择服务</span>
        <div class="top_first_list">
          <ul>
            <li> {{if .IsLogin}}
      <a href="/frame" >讨债</a>
      {{else}}
      <a href="/login" >讨债</a>
      {{end}}</li>
            <li>{{if .IsLogin}}
      <a href="/frame2" >报仇</a>
      {{else}}
      <a href="/login" >报仇</a>
      {{end}}</li>
            <li>{{if .IsLogin}}
      <a href="/frame3" >泼粪</a>
      {{else}}
      <a href="/login" >泼粪</a>
      {{end}}</li>
            <li>{{if .IsLogin}}
      <a href="/frame4" >保镖</a>
      {{else}}
      <a href="/login" >保镖</a>
      {{end}}</li>
            <li>{{if .IsLogin}}
      <a href="/frame5" >强抢民女</a>
      {{else}}
      <a href="/login" >强抢民女</a>
      {{end}}</li>
          </ul>
        </div>
      </div>
      <ul>
        <li {{if .IsHome}}class="menu"{{end}}><a href="/" style="text-decoration:none;color:red; display:block;">首页</a></li>
        <li {{if .IsMessage}}class="menu"{{end}}><a href="/message" style="text-decoration:none;color:red; display:block;">留言</a></li>
        <li {{if .IsPolice}}class="menu"{{end}}><a href="/police" style="text-decoration:none;color:red; display:block;">报警</a></li>
      </ul>
       <marquee  scrollamount=10 width=600><span style="color:blue;line-height:50px;font-size:30px; font-family:'微软雅黑'">欢迎来到<span style="color:#36dd98">浪里格朗</span>黑社会服务有限公司</span></marquee>
       {{if .IsLogin}}
      <a href="/login?exit=true" style="color:yellow;font-size:30px; line-height:50px; display:block;height:50px;width:100px;float:right; text-decoration:none;">退出</a>
      {{else}}
      <a href="/login" style="color:yellow;font-size:30px; line-height:50px; display:block;height:50px;width:100px;float:right; text-decoration:none;">登陆</a>
      {{end}}
    </div>
{{end}}