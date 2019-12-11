<template>
	<van-tabs v-model="active" color="#ff0000" title-active-color="#ff0000" @change="tabChange" swipeable animated sticky :offset-top="offsetTop">
		<van-tab v-for="(item,index) in tabs" :key="index" :title="item.Name">
			<van-list
				v-model="loading"
				:finished="finished"
				finished-text="没有更多了"
				@load="onLoad"
			>
				<double-card :colSpan="colSpan" :doubleCardItems="goods" fromPage="mall"></double-card>
			</van-list>
		</van-tab>
	</van-tabs>
</template>

<script>
	import {Tab,Tabs,List,Toast} from 'vant'
	import doubleCard from '@/components/double-card.vue'
	export default{
		components:{
			[Tab.name]:Tab,
			[Tabs.name]:Tabs,
			[List.name]:List,
			[Toast.name]:Toast,
			doubleCard
		},
		data(){
			return{
				active:0,
				tabs:[{
					ID:0,
					ParentID:0,
					Name:'全部商品',
					Status:0,
					SortOrder:0,
					Detail:'',
					CreateTime:'',
					UpdateTime:''
				}],
				colSpan:12,
				goods:[],
				loading:false,
				finished:false,
				offsetTop:46
			}
		},
		methods:{
			tabChange(name){
				this.loading = false,
				this.finished = false
				this.goods = []
				this.$store.commit('requestMallcount',0)
				var pid = this.tabs[name].ID
				this.getGoods(this.$store.state.requestMallcount,pid)
			},
			getGoods(count,pid){
				if (pid === 0) {
					var url = '/product/allList?page='+count
					this.axios.get(url).then((response)=>{
						if (response.data.status === 1) {
							if (response.data.obj.length > 0) {
								this.$store.commit('requestMallcount',1)
								this.goods = [...this.goods,...response.data.obj]
							} else{
								this.finished = true
							}
						}
					}).catch((response)=>{
						Toast('网络请求出错，请重试')
					})
				} else{
					var url = '/product/getProductList?page='+count+'&category='+pid
					this.axios.get(url).then((response)=>{
						if (response.data.status === 1) {
							if (response.data.obj.length > 0) {
								this.$store.commit('requestMallcount',1)
								this.goods = [...this.goods,...response.data.obj]
							} else{
								this.finished = true
							}
						}
					}).catch((response)=>{
						Toast('网络请求出错，请重试')
					})
				}
			},
			getCategory(){
				var url = '/product/getCat'
				this.axios.get(url).then((response)=>{
					// console.log(response)
					if (response.data.status === 1 && response.data.obj.length > 0) {
						this.tabs = [...this.tabs,...response.data.obj]
					} else{
						Toast("数据请求出错，请重试")
					}
				}).catch((response)=>{
					Toast("网络请求出错，请重试")
				})
			},
			onLoad(){
				var pid = this.tabs[this.active].ID
				setTimeout(()=>{
					this.getGoods(this.$store.state.requestMallcount,pid)
					this.loading = false
				},500)
			}
		},
		mounted() {
			this.getCategory()
			this.$store.commit('requestMallcount',0)
			this.getGoods(this.$store.state.requestMallcount,this.active)
		}
	}
</script>

<style scoped>
	.van-tab{
		padding: 0;
	}
</style>
