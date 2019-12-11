<template>
	<div>
		<van-nav-bar title="所有用户信息" fixed left-text="返回" left-arrow @click-left="onClickLeft" />
		<div class="height"></div>
		<table>
			<thead>
				<tr class="top">
					<th width="60px">ID</th>
					<th width="70px">姓名</th>
					<th width="60px">权限</th>
					<th width="70px">年龄</th>
					<th width="150px">联系电话</th>
				</tr>
			</thead>
			<tbody id="body" class="body">
			</tbody>
		</table>
	</div>
</template>

<script>
	import $ from 'jquery'
	export default {
		data(){
			return{
				html:'',
				token:''
			}
		},
		methods: {
			onClickLeft() {
				this.$router.go(-1)
			}
		},
		beforeMount: function() {
			var info = localStorage.getItem("info")
			info = JSON.parse(info)
			this.token=info.token
			$.ajax({
				url: "http://47.98.227.139:9001/user/getall",
				type: "get",
				data:{
					token:this.token
				},
				dataType: "json",
				success: (data) => {
					this.$toast(data.desc);
					if (data.desc && data.desc === 'token异常，重新登陆') {
						
						this.$router.push('/');
					}
					var json = data.data
					
					var html=''
					window.console.log(json)
					this.totals = json.length
					$.each(json, function(index, item) {
						html += "<tr>";
						html += "<td>" + item.id + "</td>";
						html += "<td>" + item.username + "</td>";
						html += "<td>" + item.is_m + "</td>";
						html += "<td>" + item.age + "</td>";
						html += "<td>" + item.tel + "</td>";
						//html += "<td>" + item.desc + "</td>";
						html += "</tr>";
					})	
					$("#body").html(html)
				},
				error: function() {
					alert("can't find server")
				}
			})
		}
	}
</script>

<style scoped="scoped">
	.height{
		margin-top: 60px;
	}
	/* .top{
		background-color: white;
		position: fixed;
		top:100px;
	} */
	.body{
		font-size: 16px;
		font-style: italic;
		line-height: 50px;
	}
</style>
