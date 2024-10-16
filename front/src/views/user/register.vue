<template>
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
        <div class="success" style="position: absolute; right:5%;top:3%" v-show="success">
            <div role="alert" class="alert alert-success">
                <svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6" fill="none"
                    viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                <span v-text="message"></span>
            </div>
        </div>
        <div class="flex items-center justify-center min-h-screen bg-gray-100">
            <div class="w-full max-w-xs p-8 bg-white rounded-3xl shadow-lg">
                <div class="mb-4 text-center">
                    <div class="text-5xl font-bold text-blue-600">HNKJZYXY</div>
                    <h2 class="text-xl font-semibold">区块链养老保险平台</h2>
                </div>
                <form class="mb-4">
                    <div class="mb-4"><label class="block mb-2 text-sm font-bold text-gray-700" for="email">
                            请输入用户名
                        </label><input
                            class="w-full px-3 py-2 leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline"
                            id="email" placeholder="用户名" type="text" v-model="username"></div>
                    <div class="mb-3"><label class="block mb-2 text-sm font-bold text-gray-700" for="password">
                            请输入密码
                        </label><input
                            class="w-full px-3 py-2 mb-3 leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline"
                            id="password" placeholder="密码" type="password" v-model="password">
                    </div>
                    <div class="mb-6"><label class="block mb-2 text-sm font-bold text-gray-700" for="password">
                            请再次输入密码
                        </label><input
                            class="w-full px-3 py-2 mb-3 leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline"
                            id="again_password" placeholder="再次输入密码" type="password" v-model="again_password">
                    </div>
                    <div class="mb-6"><button
                            class="inline-flex items-center justify-center text-sm ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 h-10 w-full bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                            @click="register">
                            注册
                        </button></div>
                </form>
                <p class="text-center text-sm">
                    已有账号?
                    <a class="text-blue-600 hover:text-blue-800" @click="gotoLogin" style="cursor: pointer;">
                        登录
                    </a>
                </p>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "register",
    data() {
        return {
            username: '',
            password: '',
            again_password: '',
            alert: false,
            success: false,
            message: ''
        }
    },
    methods: {
        gotoLogin() {
            this.$router.push('/login')
        },
        register(event) {
            // 阻止点击默认事件发生
            event.preventDefault();
            if (this.password != this.again_password) {
                this.message = "两次输入密码不相同!"
                this.alert = true
                setTimeout(() => {
                    this.alert = false;
                }, 1000); // 1秒后执行
                return
            }

            if (this.username.length < 6 || this.password.length < 6) {
                this.message = "用户名或密码长度不能短于6!"
                this.alert = true
                setTimeout(() => {
                    this.alert = false;
                }, 1000); // 1秒后执行
                return
            }

            this.$axios.post('/user/register', {
                username: this.username,
                password: this.password
            }).then((rep) => {
                if (rep.data.code != 200) {
                    this.message = rep.data.message
                    this.alert = true;
                    this.username = ''
                    this.password = ''
                    this.again_password=''
                    setTimeout(() => {
                        this.alert = false;
                    }, 1000); // 1秒后执行
                    return
                } else {
                    this.message = '注册成功！'
                    this.success = true
                    setTimeout(() => {
                        this.success = false;
                        // 跳转到登录耶main
                      this.$router.push('/login')
                    }, 1000); // 1秒后执行
                }
            }).catch((rep) => {
                this.message = rep.response.data.message
                this.alert = true;
                this.username = ''
                this.password = ''
                this.again_password=''
                setTimeout(() => {
                    this.alert = false;
                }, 1000); // 1秒后执行
            })
        }
    }
}
</script>

<style scoped></style>