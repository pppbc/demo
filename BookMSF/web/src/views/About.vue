<template>
	<div class="about">

		<!-- 头部 -->
		<div class="header">
			<img src="http://47.98.227.139:81/icon-cake.png" class="head">
			<span class="change">
				<van-icon name="setting-o" size="22px" @click="showRightTop"></van-icon>
			</span>
			<span class="name">{{localname}}</span>
			<span class="desc">{{desc}}</span>

		</div>

		<!-- 信息列表 -->
		<van-cell-group>
			<van-cell class="left" icon="records" title="ID" :value="id" />
			<van-cell class="left" icon="points" title="年龄" :value="age" />
			<van-cell class="left" icon="phone-o" title="电话号码" :value="tel" />
			<van-cell class="left" icon="gem-o" title="权限等级" :value="ism" />
			<van-cell class="left" icon="todo-list-o" title="借阅信息" :value="lend" is-link @click="listinfo" />
		</van-cell-group>

		<!-- 弹出层 -->
		<van-popup v-model="showItem" position="bottom" :style="{height:'80%'}">
			<div style="margin:20px 10px;color:gray;">更改信息</div>
			<!-- 输入框 -->
			<van-uploader v-model="fileList" multiple :after-read="afterRead" />
			<van-cell-group>
				<van-field v-model="localname" required label="用户名" right-icon="question-o" placeholder="请输入用户名" />

				<van-field v-model="age" type="text" label="年龄" placeholder="请输入年龄" />
				<van-field v-model="tel" type="tel" label="手机号码" placeholder="请输入手机号码" right-icon="question-o" />
				<van-field v-model="ism" type="text" label="权限" placeholder="请输入权限等级" disabled @click="clickUnable()" />
				<!-- 高度自适应 -->
				<van-field v-model="desc" label="签名" type="textarea" placeholder="请输入签名" rows="1" autosize />
			</van-cell-group>

			<input class="submit" @click="submit()" type="submit" value="确认修改">
		</van-popup>

		<!-- 我的借书信息 -->
		<van-popup v-model="isShow" position="bottom" :style="{height:'50%'}">
			<div id="num">共:<span class="num">{{num}}</span> 条借书信息</div>
			<table>
				<thead>
					<tr class="thead">
						<th width='100px'>书名</th>
						<th width='100px'>借书时间</th>
						<th width='50px'>状态</th>
						<th width='100px'>还书时间</th>
					</tr>
				</thead>
				<tbody id="showbookinfo" class="tbody"></tbody>
			</table>

		</van-popup>

	</div>
</template>

<script>
	import $ from 'jquery'
	import {
		Row,
		Col,
		Icon,
		Cell,
		CellGroup,
		Field,
		Uploader,
		Toast
	} from 'vant';
	export default {
		data() {
			return {
				num: 0,
				isShow: false,
				showItem: false,
				localname: '未登录',
				// pullItems:Array,
				id: 0,
				age: 0,
				tel: '',
				ism: 0,
				lend: '',
				desc: '这个人很懒，什么都没有说',
				fileList: [{
					url: 'http://47.98.227.139:81/icon-cake.png'
				}, ],
				token:''
			}
		},
		components: {
			[Row.name]: Row,
			[Col.name]: Col,
			[Icon.name]: Icon,
			[Cell.name]: Cell,
			[CellGroup.name]: CellGroup,
			[Field.name]: Field,
			[Uploader.name]: Uploader,
			[Toast.name]: Toast
		},
		beforeMount: function() {
			//渲染前获取登陆时本地存储的用户名
			var info = localStorage.getItem("info")
			info = JSON.parse(info)
			this.localname = info.username
			this.id = info.id
			this.age = info.age
			this.tel = info.tel
			this.ism = info.root
			this.desc = info.desc
			this.token=info.token

			// $.ajax({
			// 	url: "http://47.98.227.139:9001/user/get",
			// 	type: "get",
			// 	data: {
			// 		admin: this.localname,
			// 		token:this.token
			// 	},
			// 	dataType: "json",
			// 	success: (data) => {
			// 		var json = data.data
			// 		window.console.log(json)
			// 		this.id = json.id
			// 		this.age = json.age
			// 		this.tel = json.tel
			// 		this.ism = json.is_m
			// 		this.desc = json.desc
			// 	},
			// 	error: function() {
			// 		alert("can't find server")
			// 	}
			// })
		},
		methods: {
			listinfo() {
				this.isShow = true

				$.ajax({
					url: "http://47.98.227.139:9001/lend/getbyuser",
					type: "get",
					dataType: "json",
					data: {
						username: this.localname,
						token:this.token
					},
					success: (data) => {
						if (data.desc && data.desc === 'token异常，重新登陆') {
							this.$toast(data.desc);
							this.$router.push('/');
						}
						
						var json = data.data
						if (json) {
							this.num = json.length
							if (json) {
								if (json) {
									var html = ''
									$.each(json, function(index, item) {
										html += "<tr>";
										html += "<td>" + item.bookname + "</td>";
										html += "<td>" + item.lend_time + "</td>";
										html += "<td>" + item.id_r + "</td>";
										html += "<td>" + item.return_time + "</td>";
										html += "</tr>";
									})
									html += "</tbody>"
									$("#showbookinfo").html(html)
									html = ''
								} else {

								}
							}
						}
					},
					error: function() {
						alert("can't find server")
					}
				})
			},
			showRightTop() {
				this.showItem = true
			},
			afterRead(file) {
				// 此时可以自行将文件上传至服务器
				window.console.log(file);
			},
			clickUnable() {
				this.$toast('无权限！');
			},
			submit() {
				window.console.log("提交拉")
				$.ajax({
					url: "http://47.98.227.139:9001/user/change",
					type: "post",
					data: {
						id: this.id,
						username: this.localname,
						age: this.age,
						tel: this.tel,
						is_m: this.ism,
						desc: this.desc,
						token:this.token
					},
					dataType: "json",
					success: (data) => {
					
					if (data.desc && data.desc === 'token异常，重新登陆') {
						this.$toast(data.desc);
						this.$router.push('/');
					}
						this.$toast(data.desc)
						//获取返回数据
						var json = data.data
						window.console.log(json)
						this.localname = json.username
						this.id = json.id
						this.age = json.age
						this.tel = json.tel
						this.ism = json.is_m
						this.desc = json.desc

						//修改成功则存储用户名在本地
						var info = {
							username: json.username,
							id:json.id,
							age:json.age,
							head_img:json.head_img,
							tel:json.tel,
							desc:json.desc,
							token:json.token,
							root:json.is_m
						}
						info = JSON.stringify(info)
						localStorage.setItem("info", info)
					},
					error: () => {
						this.$toast('can‘t find serve！')
						// alert("can't find server")
					}
				})
				this.showItem = false
			}
		}
	}
</script>

<style scoped="scoped">
	.header {
		background-color: mediumslateblue;
		padding: 20px 0 10px 0;
	}

	.head {
		display: block;
		width: 120px;
		height: 120px;
		border-radius: 50%;
		margin-left: 5px;
	}

	.name {
		position: absolute;
		top: 70px;
		left: 140px;
		font-size: 18px;
		color: white;
		font-weight: bolder;
	}

	.desc {
		position: absolute;
		top: 100px;
		left: 140px;
		font-size: 14px;
		color: white;
	}

	.change {
		position: absolute;
		top: 20px;
		right: 20px;
		font-size: 14px;
		color: black;
	}

	.left {
		text-align: left;
	}

	.content {
		margin: 50px 0;
	}


	.user-group {
		margin-bottom: 15px;
	}

	.user-links {
		padding: 15px 0;
		font-size: 12px;
		text-align: center;
		background-color: #fff;
	}

	.user-links>.van-icon {
		display: block;
		font-size: 24px;
	}

	.submit {
		margin-top: 50px;
		/* border: 1px solid orangered; */
		border-radius: 5px;
		font-size: 16px;
		background-color: whitesmoke;
	}

	.thead {
		/* color:#F08080 */
		background-color: floralwhite;
	}

	.tbody {
		color: #F08080;
		font-size: 14px;
		line-height: 30px;
		font-style: italic;
	}

	#num {
		margin: 20px 0;
	}

	.num {
		color: lightcoral;
		font-style: italic;
	}
</style>
