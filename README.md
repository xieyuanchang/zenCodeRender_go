# zenCodeRender_go 
Zen Code 缩写扩展器(go语言实现版本)

以CSS缩写形式对Zen Code进行HTML扩展
 
# 外部链接参考
Zen Code 背景 参考百度百科：http://baike.baidu.com/view/3371595.htm

Emmet Documentation http://docs.emmet.io/

# Zen Code 编写实例
	Zen Code：
		div#page>div.logo+ul#navigation>li*5>a
	
	缩写扩展器处理后：
		<div id="page">
		<div class="logo"></div>
		<ul id="navigation">
			<li><a></a></li>
			<li><a></a></li>
			<li><a></a></li>
			<li><a></a></li>
			<li><a></a></li>
		</ul>
		</div>
