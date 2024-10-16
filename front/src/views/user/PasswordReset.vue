<template>
    <!-- 拟态框 -->
    <div class="model">
        <div class="modal" role="dialog" id="my_modal_10">
            <div class="modal-box">
                <h3 class="font-bold text-lg">用户信息</h3>
                <div class="username" style="margin-top: 4%;margin-left:18%">
                    <span style="font-weight: bolder;">用户名: </span>
                    <span style="margin-left:10%" v-text="username"></span>
                </div>
                <div class="password" style="margin-top: 4%;margin-left:18%">
                    <span style="font-weight: bolder;">密&nbsp;&nbsp;&nbsp;&nbsp;码: </span>
                    <span style="margin-left:10%;" v-text="password"></span>
                </div>
                <p class="py-4" style="font-size: 10px;font-weight:bolder;position:absolute;bottom:4%">注意: 请登录自行修改密码</p>
                <div class="modal-action">
                    <a href="#" class="btn">关闭</a>
                </div>
            </div>
        </div>
    </div>
    <div class="app">
        <div class="success" style="position: absolute; right:5%;top:3%" v-show="alert">
            <div role="alert" class="alert alert-error">
                <svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6" fill="none"
                    viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                <span v-text="message"></span>
            </div>
        </div>
        <div class="flex justify-center items-center h-screen bg-gray-100">
            <div class="w-full max-w-xs mx-auto">
                <div class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
                    <div class="mb-4">
                        <h1 class="block text-blue-500 font-bold mb-2 text-xl text-center">HNKJZYXY</h1>
                        <p class="text-gray-700 text-base mb-4 text-center">区块链养老保险平台</p>
                    </div>
                    <form>
                        <div class="mb-4"><label class="block text-gray-700 text-sm font-bold mb-2" for="email">
                                姓名
                            </label><input
                                class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                                id="email" placeholder="真实姓名" type="text" v-model="name"></div>
                        <div class="mb-6"><label class="block text-gray-700 text-sm font-bold mb-2" for="password">
                                身份证
                            </label><input
                                class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline"
                                id="password" placeholder="身份证" type="text" v-model="id_card">
                        </div>
                        <div class="flex items-center justify-between"><button
                                class="inline-flex items-center justify-center text-sm ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 h-10 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                                @click="FindPassword">找回密码</button>
                        </div>
                        <div class="mt-4 text-center">
                            <p class="text-gray-600 text-xs">
                                没有绑定身份信息? <a class="text-blue-500 hover:text-blue-800" style="cursor:pointer"
                                    @click="gotoRegister">注册</a>
                            </p>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "PasswordReset",
    data() {
        return {
            name: "", // 姓名
            id_card: "", // 身份证
            alert: false,
            success: false,
            message: '',
            username: '',
            password: ''
        }
    },
    methods: {
        // 跳转到注册页面
        gotoRegister() {
            this.$router.push('/register')
        },
        // 找回密码
        FindPassword(event) {
            // 阻止点击默认事件发生
            event.preventDefault();
            if (this.name.length < 0 || this.id_card.length != 18) {
                this.message = "请认真填写用户信息！！！"
                this.alert = true
                setTimeout(() => {
                    this.alert = false;
                }, 1000); // 1秒后执行
                return
            }
            this.$axios.post("/user/findPassword", {
                name: this.name,
                idCard: this.id_card
            }).then((rep) => {
                if (rep.data.code != 200) {
                    this.message = rep.data.message
                    this.alert = true
                    setTimeout(() => {
                        this.alert = false;
                    }, 1000); // 1秒后执行
                    return
                }
                this.username = rep.data.username
                this.password = rep.data.password
                window.location.href = '#my_modal_10'
            }).catch((rep) => {
                if (rep.code != 200) {
                    this.message = rep.response.data.message
                    this.alert = true
                    setTimeout(() => {
                        this.alert = false;
                    }, 1000); // 1秒后执行
                    return
                }
            })
        }
    }
}
</script>

<style scoped></style>