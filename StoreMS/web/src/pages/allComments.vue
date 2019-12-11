<template>
	<div>
		<van-nav-bar
			title="全部评价"
			left-arrow
			@click-left="goBack"
		/>
		
		<van-list
			v-model="loading"
			:finished="finished"
			finished-text="没有更多了"
			@load="onLoad"
		>
			<van-row class="comment">
				<van-col span="24" v-for="(item,index) in allComments" :key="index">
					<van-row class="name">
						<van-col span="10" offset="1">{{item.UserName.replace(item.UserName.substring(3,7),'****')}}</van-col>
						<van-col span="5">{{item.Title}}</van-col>
					</van-row>
					<van-row class="time">
						<van-col span="23" offset="1">{{item.CreateTime.replace('T',' ').replace('Z',' ')}}</van-col>
					</van-row>
					<van-row class="content">
						<van-col span="23" offset="1">{{item.Content}}</van-col>
					</van-row>
				</van-col>
			</van-row>
		</van-list>
		
	</div>
</template>

<script>
	import {NavBar,Row,Col,List,Toast} from 'vant'
	export default{
		components:{
			[NavBar.name]:NavBar,
			[Row.name]:Row,
			[Col.name]:Col,
			[List.name]:List,
			[Toast.name]:Toast
		},
		data(){
			return{
				allComments:[],
				loading:false,
				finished:false,
				pid:0
			}
		},
		methods:{
			goBack(){
				this.$router.back()
			},
			onLoad(){
				setTimeout(()=>{
					this.getAllComments(this.$store.state.requestAllCommentsCount)
					this.loading = false
				},100)
			},
			getAllComments(count){
				this.$store.commit('allComments','')
				var url = '/order/getProductComment?product_id='+this.pid+'&page='+count
				this.axios.get(url).then((response)=>{
					if (response.data.status === 1) {
						if (response.data.obj.length > 0) {
							this.$store.commit('requestAllCommentsCount',1)
							this.$store.commit("allComments",response.data.obj)
							this.allComments = [...this.allComments,...response.data.obj]
						} else{
							this.finished = true
						}
					} 
				}).catch((response)=>{
					Toast("网络请求出错")
				})
			}
		},
		beforeMount() {
			this.$store.commit('requestAllCommentsCount',0)
			this.$store.commit("allComments",'')
			this.pid = this.$route.params.pid
		}
	}
</script>

<style scoped>
	.van-nav-bar__arrow{
		font-size: 1.5rem;
		color: black;
	}
	
	.comment{
		background: white;
	}
	.van-col{
		margin-top: 0.2rem;
		margin-bottom: 0.2rem;
	}
	.comment>.van-col{
		margin-bottom: 0.3125rem;
	}
	.name{
		text-align: left;
		font-size: 0.875rem;
	}
	.time{
		font-size: 0.75rem;
		color: gray;
		text-align: left;
	}
	.content{
		font-size: 0.9375rem;
		text-align: left;
	}
</style>
