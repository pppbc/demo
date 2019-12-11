<template>
	<div class="justlook">
		<!-- 返回栏目 -->
		<van-nav-bar title="书籍列表" fixed left-text="返回" left-arrow @click-left="onClickLeft" />

		<!-- cell单元格，提示框 -->
		<van-cell class="header" title-style="font-weight:bold" title="古典文学,国粹经典" :value="totals" is-link label="相关书籍" />
		<!-- 瀑布流 -->

		<div class="root">
			<div class="item" v-for="(item,index) in pullItems" :key="index">
				<img class="itemImg" :src="item.src" alt="#">
				<div class="userInfo">
					<span class="author">{{item.author}} <em>剩:{{item.num}}本</em></span><br>
					<span class="title">{{item.title}}</span><br>
				</div>
			</div>
		</div>

		<!-- 第二个书籍种类 -->
		<van-cell class="footer" title-style="font-weight:bold" title="外语文学,世界经典" :value="totals" is-link label="相关书籍" />
		<div class="root">
			<div class="item" v-for="(item,index) in pullItems" :key="index">
				<img class="itemImg" :src="item.src" alt="#">
				<div class="userInfo">
					<img class="avater" src="../assets/images/icon-heart.png" alt="#">
					<span class="author">{{item.author}} <em>{{item.num}}</em></span><br>
					<span class="title">{{item.title}}</span><br>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
	import $ from 'jquery'
	import {
		Card,
		Cell
	} from 'vant'
	export default {
		data() {
			return {
				html: '',
				totals: 0,
				pullItems: Array
			}
		},
		components: {
			[Card.name]: Card,
			[Cell.name]: Cell,
		},
		beforeMount: function() {
			var info = localStorage.getItem("info")
			info = JSON.parse(info)
			this.token=info.token
			$.ajax({
				url: "http://47.98.227.139:9001/book/getall",
				type: "get",
				data:{
					
				},
				dataType: "json",
				success: (data) => {
					var json = data.data
					window.console.log(json)
					window.console.log(data.status)
					window.console.log(data.desc)
					this.totals = json.length
					$("#total").html(this.totals)
					this.pullItems = json
					window.console.log(json)
					window.console.log(this.pullItems)
				},
				error: function() {
					alert("can't find server")
				}
			})
		},
		mounted: function() {
			// window.console.log(this.pullItems)
		},
		methods: {
			onClickLeft() {
				this.$router.go(-1)
			}
		}
	}
</script>

<style scoped="scoped">
	body {
		background: #e5e5e5;
	}

	/*  */
	.header,
	.footer {
		text-align: left;
		background-color: #FaFfFf;
	}

	.header {
		margin-top: 45px;
	}

	/* 瀑布流 */
	.root {
		column-count: 2;
		column-width: auto;
		margin: 0 10px;
		/* background-color: #DCDCDC; */
		/* border: 1px solid gold; */
	}

	/*  每一列图片包含层 */
	.item {
		margin: 0 0 1rem 0;
		/* 防止多列布局,分页媒体和多区域上下文中的意外中断 */
		break-inside: avoid;
		background: #fff;
		border-radius: 8px;
		/* margin:1rem 0 0 0; */
	}

	.item:hover {
		box-shadow: 2px 2px 2px rgba(0, 0, 0, 0.5);
	}

	.itemImg {
		width: 100%;
		vertical-align: middle;
		border-radius: 8px;
	}

	/* 图片下的信息包含层 */
	.userInfo {
		padding: 5px 10px;
	}

	.avater {
		vertical-align: middle;
		width: 30px;
		height: 30px;
		border-radius: 50%;
	}

	.title,
	.author {
		font-size: small;
		margin-left: 5px;
		text-shadow: 1px 1px 1px rgba(0, 0, 0, .3);
		color: lightslategray;
	}

	.author>em {
		color: #7B68EE
	}

	.title {
		font-size: 14px;
		color: #555;
	}
</style>
