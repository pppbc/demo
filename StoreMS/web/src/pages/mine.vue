<template>
	<div id="mine">
		<img class="bg" :src="bg" alt="">
		<van-button class="settingBtn" @click="goSettings"/>
		<van-nav-bar 
			title="我的"
			:border="border"
			fixed
		>
			<van-icon name="chat-o" slot="right" @click="goChatPage" />
		</van-nav-bar>
		
		<van-row class="header" @click="myInfo">
			<van-col span="8" class="header_left">
				<van-image :src="getAvatar" width="5rem" height="5rem" radius="0.3125rem"/>
			</van-col>
			<van-col span="12" class="header_center">
				<div>
					<div>{{nickname}}</div>
					<div>{{username}}</div>
				</div>
			</van-col>
			<van-col span="4" class="header_right">
				<van-icon :name="arrow" size="1.5rem"/>
			</van-col>
		</van-row>
		
		<div class="content">
			<van-cell-group v-for="(datas,index) in dataList" :key="index">
				<van-cell v-for="(data,index) in datas" :key="index" 
					:icon="data.image"
					:title="data.title"
					:value="hasValue(data.integral)"
					:to="data.name"
					is-link
				>
				</van-cell>
			</van-cell-group>
		</div>
		
	</div>
</template>

<script>
	import {NavBar,Toast,CellGroup,Cell,Button,Row,Col,Image,Icon} from 'vant'
	import {Images} from '../resources/index.js'
	export default{
		components:{
			[NavBar.name]:NavBar,
			[Toast.name]:Toast,
			[CellGroup.name]:CellGroup,
			[Cell.name]:Cell,
			[Button.name]:Button,
			[Row.name]:Row,
			[Col.name]:Col,
			[Image.name]:Image,
			[Icon.name]:Icon
		},
		data(){
			return{
				bg:Images.common.ic_header_bg,
				border:false,
				arrow:Images.mineView.ic_go,
				dataList:[
					[
						{image:Images.mineView.ic_label_data[0],title:'我的积分',integral:this.$store.state.totalBlock,name:'/wallet'},
					],
					[
						{image:Images.mineView.ic_label_data[1],title:'待付款',name:'/orderPage/unpaid/待付款'},
						{image:Images.mineView.ic_label_data[2],title:'待发货',name:'/orderPage/unshipped/待发货'},
						{image:Images.mineView.ic_label_data[3],title:'待收货',name:'/orderPage/unreceived/待收货'},
						{image:Images.mineView.ic_label_data[4],title:'待评价',name:'/orderPage/comment/待评价'},
						{image:Images.mineView.ic_label_data[5],title:'全部订单',name:'/orderPage/allorders/全部订单'},
						{image:Images.mineView.ic_label_data[6],title:'退货/退款',name:'/refundAndReturnPage/退款'},
					],
					[
						{image:Images.mineView.ic_label_data[7],title:'问题反馈',name:'/feedback'},
						{image:Images.mineView.ic_label_data[8],title:'关于我们',name:''}
					]
				]
			}
		},
		methods:{
			goSettings(){
				this.$router.push('/setting')
			},
			myInfo(){
				this.$router.push('/userinfo')
			},
			goChatPage(){
				this.$store.commit('fromPage','mine')
				this.$router.push('/chatPage')
			}
		},
		computed:{
			getAvatar(){
				return this.$store.state.avatar ? this.$store.state.avatar : Images.mineView.ic_default_head1
			},
			nickname(){
				return this.$store.state.nickname
			},
			username(){
				return this.$store.state.username.replace(/(\d{3})\d{4}(\d{4})/,'$1****$2')
			},
			hasValue(){
				return function(value){
					return value ? value + '积分' : ''
				}	
			}
		}
	}
</script>

<style scoped>
	.bg{
		position: fixed;
		top: 0;
		left: 0;
		width: 100%;
		height: 7.5rem;
		z-index: -1;
	}
	.settingBtn{
		position: fixed;
		top: 0.625rem;
		left: 0.625rem;
		margin: 0;
		padding: 0;
		border: none;
		width: 1.5rem;
		height: 1.5rem;
		background: url(../assets/Images/Mine/set.png) no-repeat center;
		background-size: 100% 100%;
		z-index: 2;
	}
	.van-nav-bar{
		background: url(../assets/Images/Home/header_bg.png);
	}
	.van-nav-bar__title{
		color: #fff;
		font-size: 1.125rem;
	}
	.van-nav-bar .van-icon{
		color: #FFFFFF;
		font-size: 1.5rem;
	}
	.header{
		background-color: white;
		height: 6.25rem;
		line-height: 6.25rem;
		margin: 0.625rem 0.625rem 0 0.625rem;
		border-radius: 0.5rem;
	}
	.header:active{
		background: #f2f3f5;
	}
	.header .van-col{
		height: 6.25rem;
		line-height: 6.25rem;
	}
	.header .header_left .van-image{
		margin-top: 0.625rem;
	}
	.header .header_center>div{
		text-align: left;
		margin: 1rem 0;
		padding: 0;
		height: 4.75rem;
		line-height: 2.35rem;
	}
	.header .header_right{
		padding-top: 0.625rem;
	}
	.content{
		margin: 0.625rem 0.625rem 0 0.625rem;
		border-radius: 0.5rem;
	}
	.content .van-cell-group{
		margin-bottom: 0.3125rem;
		border-radius: 0.5rem;
		overflow: hidden;
	}
	.content .van-cell-group .van-cell{
		/* border-radius: 0.3125rem; */
		text-align: left;
	}
</style>
