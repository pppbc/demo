<template>
	<div class="home">
		<van-search placeholder="搜索书名" v-model="menu" show-action @search="onSearch">
			<!-- slot="action"自定义插槽 -->
			<van-icon slot="action" name="user-o" @click="showRightTop">{{localname}}</van-icon>
		</van-search>

		<!-- 功能版块 -->
		<grid :gridItems="gridItems" :columnNum="columnNum" :showClass="showClass" :gutter="gutter"></grid>

		<!--  轮播-->
		<!-- <swipe :swipeItems="swipeItems" :autoplay='autoplay' class="swipeImgHome"></swipe> -->


		<!-- 切换账户 -->
		<van-popup class="showbtn" v-model="showItem" position="top" :style="{height:'20%'}">
			<div>
				切换用户 <br>
				<van-icon class="item" name="add-o" size="26px" color="mediumslateblue" @click="changename()"></van-icon>
			</div>
			<div>
				修改密码 <br>
				<van-icon class="item" name="manager" size="26px" color="mediumslateblue" @click="changepassword()"></van-icon>
			</div>
		</van-popup>

		<!-- 搜索弹出框 -->
		<van-popup v-model="isShow" position="bottom" :style="{height:'40%'}" @click-overlay="clickoverlay">
			<van-cell-group>
				<van-field v-model="title" label="书名" placeholder="书名" disabled />
				<van-field v-model="author" label="作者" placeholder="作者" disabled />
				<van-field v-model="bookid" label="编号" placeholder="编号" disabled />
				<van-field v-model="num" type="number" label="剩余数量" placeholder="余量" disabled />
				<van-field v-model="price" type="number" label="价格" placeholder="价格" disabled />
				<van-field v-model="desc" label="书籍详情" type="textarea" placeholder="请输入书籍描述" rows="2" autosize disabled />
			</van-cell-group>
		</van-popup>
		<div style="height: 60px;">
		</div>
	</div>
</template>

<script>
	import {
		Search,
		Icon,
		Popup
	} from 'vant'
	import $ from 'jquery'
	import swipe from '@/components/swipe.vue'
	import grid from '@/components/grid.vue'
	export default {
		components: {
			grid,
			swipe,
			[Search.name]: Search,
			[Icon.name]: Icon,
			[Popup.name]: Popup
		},
		data() {
			return {
				autoplay: 3000,
				menu: 'vue从入门到放弃',
				localname: '未登陆',
				showItem: false,
				gutter: 0,
				columnNum: 1,
				showClass: 'gridHome',
				root: 0,
				swipeItems: [
					require('../assets/images/bg-cake.jpg'),
					require('../assets/images/bg-panda.jpg'),
					require('../assets/images/bg-meat.jpg')
				],
				gridItems: [{
						text: '查看所有书籍',
						path: '/justlook',
						icon: 'description'
					},
					{
						text: '功能6',
						path: '/haha',
						icon: 'circle'
					},
					{
						text: '管理书籍入口',
						path: '/managebook',
						show: '',
						icon: 'add-o'
					},
					{
						text: '借书还书入口',
						path: '/lendreturn',
						show: '',
						icon: 'exchange'
					},
					{
						text: '查看用户信息',
						path: '/allusers',
						show: '',
						icon: 'user-o'
					},
					{
						text: '管理用户入口',
						path: '/manageusers',
						show: '',
						icon: 'setting-o'
					},
				],
				//
				isShow: false,
				title: '',
				author: '',
				num: 0,
				price: 0.00,
				bookname: '',
				desc: '',
				bookid: 0,
				booksrc: '',
				token:''
			}
		},
		beforeMount: function() {
			
			window.console.log("Home beforeMount")
			//渲染前获取本地存储的用户名
			var info = localStorage.getItem("info")
			info = JSON.parse(info)
			this.token=info.token
			this.localname = info.username
			this.root = info.root

			//通过权限判断显示功能版块
			if (this.root == 1) {
				// $.each(this.gridItems ,function(index,item){
				// 	item.show='IsShow'
				// })
				this.gridItems[2].show = "IsShow"
				this.gridItems[3].show = "IsShow"
				this.gridItems[4].show = "IsShow"
				this.gridItems[5].show = "IsShow"
			} else {
				if (this.root == 2) {
					this.gridItems[4].show = "IsShow"
					this.gridItems[5].show = "IsShow"
				} else {
					if (this.root == 3) {
						this.gridItems[3].show = "IsShow"
						this.gridItems[2].show = "IsShow"
					}
				}
			}

			$.ajax({
				url: "http://47.98.227.139:9001/user/get",
				type: "get",
				data: {
					username: this.localname,
					token:this.token
				},
				dataType: "json",
				success: (data) => {
						if (data.desc && data.desc === 'token异常，重新登陆') {
							this.$toast(data.desc);
							this.$router.push('/');
						}
				},
				error: () => {
					this.$toast('can‘t find serve！')
				}
			})
		},
		methods: {
			showRightTop() {
				this.showItem = true
			},
			onSearch() {
				//搜索
				window.console.log(`search:${this.menu}`)
				this.isShow = true
				//
				$.ajax({
					url: "http://47.98.227.139:9001/book/get",
					type: "get",
					data: {
						title: this.menu,
						token:this.token
					},
					dataType: "json",
					success: (data) => {
						var json = data.data

						if (json) {
							this.title = json.title
							this.author = json.author
							this.price = json.price
							this.num = json.num
							this.desc = json.desc
							this.bookid = json.id
							this.booksrc = json.src
						}
						window.console.log(json)
						this.$toast(data.msg)
					},
					error: function() {
						alert("can't find server")
					},
				})
			},
			changename() {
				this.$router.push('/')
			},
			changepassword() {
				this.$router.push('/changepassword')
			},
			clickoverlay() {
				this.isShow=false
				this.title = ''
				this.author = ''
				this.price = ''
				this.num = ''
				this.desc = ''
				this.bookid = ''
				this.booksrc = ''
			},
		}
	}
</script>

<style scoped="scoped">
	/* 切换用户 */
	.showbtn>div {
		font-size: 14px;
		width: 49%;
		display: inline-block;
		margin-top: 30px;
	}

	.showbtn>div:last-child {
		border-left: 1px solid #555555;
	}

	.item {
		margin-top: 20px;
	}

	/* 轮播 */
	.swipeImgHome {
		margin-top: 50px;
		margin-left: 1rem;
		margin-right: 1rem;
		border-radius: 0.5rem;
	}

	.swipeImgHome img {
		width: 100%;
		height: 250px;
	}
</style>
