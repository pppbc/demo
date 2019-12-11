<template>
	<div class="comment">
		<van-list
			v-model="loading"
			:finished="finished"
			finished-text="没有更多了"
			@load="onLoad"
		>
			<order-cell 
				v-for="(item,index) in data" 
				:key="index"
				title="待评价"
				:time="item.CreateTime.replace('T',' ').replace('Z',' ')"
				:image="item.ProductImage"
				:name="item.ProductName"
				:goodsNum="item.Quantity"
				:totalPrice="item.TotalPrice"
				:orderNo="item.OrderNo"
				@returnOrder="returnOrder(item)"
				@comment="comment(item)"
			></order-cell>
		</van-list>
	</div>
</template>

<script>
	import {Toast,List} from 'vant'
	import orderCell from '@/components/orderCell.vue'
	export default{
		components:{
			[Toast.name]:Toast,
			[List.name]:List,
			orderCell
		},
		data(){
			return {
				data:[],
				loading:false,
				finished:false,
			}
		},
		methods:{
			getCommentList(count){
				var url = '/order/successList?user_id='+this.$store.state.userID+'&page='+count+'&token='+this.$store.state.token
				this.axios.get(url).then((response)=>{
					if(response.data.detail && response.data.detail === 'token异常，重新登陆') {
						Toast('登陆过期，请重新登录')
						this.$router.push('/login')
					}else if (response.data.status === 1) {
						if (response.data.obj.length>0) {
							this.$store.commit('requestCommentListCount',1)
							this.data = [...this.data,...response.data.obj]
						} else{
							this.finished = true
						}
					} else{
						Toast('数据请求失败，请重试')
					}
				}).catch((response)=>{
					Toast('网络请求出错，请重试')
				})
			},
			onLoad(){
				setTimeout(()=>{
					this.getCommentList(this.$store.state.requestCommentListCount)
					this.loading = false
				},1000)
			},
			returnOrder(item){
				this.$store.commit('refundReturnData',item)
				this.$router.push('/returnApply/退货申请')
			},
			comment(item){
				this.$store.commit('commentData',item)
				this.$router.push('/commentPage')
			}
		},
		beforeRouteEnter(to,from,next) {
			next(vm=>{
				vm.$store.commit('requestCommentListCount',0)
				vm.getCommentList(vm.$store.state.requestCommentListCount)
			})
		}
	}
</script>

<style scoped>
	.comment{
		margin-top: 7.5rem;
		
	}
</style>
