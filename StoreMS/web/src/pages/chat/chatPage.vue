<template>
	<div>
		<van-nav-bar 
			title="消息" 
			left-arrow
			fixed
			@click-left="goBack"
		>
			<van-icon name="wap-home" slot="right" @click="goHome"/>
		</van-nav-bar>
		
		<div class="list">
			<van-swipe-cell v-for="(item,index) in data" :key="index">
				<van-row class="content"  @click="goChat(item)">
					<van-col span="5" class="left">
						<van-image :src="item.icon"/>
					</van-col>
					<van-col span="19" class="right">
						<van-row class="title">
							<van-col span="24">{{item.title}}</van-col>
						</van-row>
						<van-row class="msg">
							<van-col span="24">{{item.msg}}</van-col>
						</van-row>
					</van-col>
				</van-row>
				<template slot="right">
					<van-button square type="danger" text="删除" @click="deleteChat(item.id,index)"/>
				</template>
			</van-swipe-cell>
		</div>
		
		
	</div>
</template>

<script>
	import {NavBar,Icon,SwipeCell,Row,Col,Button,Image} from 'vant'
	import {Images} from '@/resources/index.js'
	export default{
		components:{
			[NavBar.name]:NavBar,
			[Icon.name]:Icon,
			[Row.name]:Row,
			[Col.name]:Col,
			[SwipeCell.name]:SwipeCell,
			[Button.name]:Button,
			[Image.name]:Image
		},
		data(){
			return{
				data:[
					{
						id:1,
						icon:this.$store.state.avatar ? this.$store.state.avatar : Images.mineView.ic_default_head1,
						title:'张三淘宝旗舰店',
						msg:'你好呢'
					},
					{
						id:2,
						icon:this.$store.state.avatar ? this.$store.state.avatar : Images.mineView.ic_default_head1,
						title:'李四',
						msg:'哈哈哈'
					},
					{
						id:3,
						icon:this.$store.state.avatar ? this.$store.state.avatar : Images.mineView.ic_default_head1,
						title:'王五',
						msg:'本次acid'
					},
					{
						id:4,
						icon:this.$store.state.avatar ? this.$store.state.avatar : Images.mineView.ic_default_head1,
						title:'赵六',
						msg:'动次打次'
					}
				]
			}
		},
		methods:{
			goBack(){
				this.$store.commit('changeTabbarActive',this.$store.state.fromPage)
				this.$router.back()
			},
			goHome(){
				this.$router.push('/main/home')
			},
			deleteChat(id,index){
				this.data.splice(index,1)
			},
			goChat(item){
				this.$store.commit('chatting',item)
				this.$router.push('/chatRoom')
			}
		}
	}
</script>

<style scoped>
	.van-nav-bar{
		background: #FF7256;;
	}
	.van-nav-bar__arrow{
		font-size: 1.5rem;
		color: #FFFFFF;
	}
	.van-nav-bar__title{
		font-size: 1.125rem;
		color: #FFFFFF;
	}
	.van-nav-bar .van-icon{
		font-size: 1.5rem;
		color: #FFFFFF;
	}
	.list{
		margin-top: 2.875rem;
		overflow: hidden;
	}
	.van-swipe-cell{
		background: #FFFFFF;
		border-bottom: 0.0625rem solid #F0F0F0;
	}
	.content{
		height: 4rem;
		line-height: 4rem;
	}
	.left .van-image{
		width: 3rem;
		height: 3rem;
		margin: 0 auto;
		vertical-align: middle;
	}
	.right{
		padding: 0.5rem 0.625rem;
		text-align: left;
	}
	.title{
		height: 1.5rem;
		line-height: 1.5rem;
	}
	.msg{
		height: 1.5rem;
		line-height: 1.5rem;
		font-size: 0.875rem;
		color: gray;
	}
	.van-button{
		height: 4rem;
		line-height: 4rem;
	}
</style>
