<template>
	<div class="root">
		<div class="item" v-for="(item,index) in items" :key="index">
			<img class="itemImg" :src="item.src" alt="#">
			<div class="userInfo">
				<img class="avater" src="../assets/images/icon-heart.png" alt="#">
				<span class="username">{{item.tips}}</span>
			</div>
		</div>
	</div>
</template>

<script>
	export default{
		props:{
			items:Array
		}
	}
</script>

<style scoped>
	body{
		background: #e5e5e5;
	}
	/*  瀑布流最外层 */
	
	.root{
		column-count: 2;
		column-width: auto;
		column-gap: 10px;
	}
	/*  每一列图片包含层 */
	.item{
		margin-bottom: 1rem;
		/* 防止多列布局,分页媒体和多区域上下文中的意外中断 */
		break-inside: avoid;
		background: #fff;
		border-radius: 8px;
	}
	.item:hover{
		box-shadow: 2px 2px 2px rgba(0,0,0,0.5);
	}
	.itemImg{
		width: 100%;
		vertical-align: middle;
		border-radius: 8px;
	}
	/* 图片下的信息包含层 */
	.userInfo{
		padding: 5px 10px;
	}
	.avater{
		vertical-align: middle;
		width: 30px;
		height: 30px;
		border-radius: 50%;
	}
	.username{
		font-size: small;
		margin-left: 5px;
		color: #555;
		text-shadow: 1px 1px 1px rgba(0,0,0,.3);
	}
</style>

