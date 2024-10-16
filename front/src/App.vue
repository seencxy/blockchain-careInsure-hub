<template>
  <div class="modal" role="dialog" id="updateUserInfo">
    <div class="flex items-center justify-center h-screen ">
      <div class="rounded-lg border bg-card text-card-foreground shadow-sm" data-v0-t="card"
           style="background-color: #fff;">
        <div class="p-6">
          <div class="space-y-8">
            <div class="space-y-2">
              <h2 class="text-3xl font-semibold">修改用户信息</h2>
            </div>
            <div class="space-y-4">
              <div class="grid grid-cols-2 gap-4">
                <div class="space-y-2"><label
                    class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                    for="first-name">用户名</label><input
                    class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                    id="first-name" placeholder="用户名" disabled v-model="username"></div>
                <div class="space-y-2"><label
                    class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                    for="last-name">密码</label><input
                    class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                    id="last-name" placeholder="密码(不修改则不填)" v-model="password"></div>
              </div>
              <div class="space-y-2"><label
                  class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                  for="email">年龄</label><input type="number"
                                                 class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                                                 id="email" placeholder="年龄" disabled v-model="age"></div>
              <div class="space-y-2"><label
                  class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                  for="email">性别</label>
                <div class="radio-group">
                  <label>
                    <input type="radio" name="gender" value="male" class="radio radio-primary" v-model="selectedGender"
                           :disabled="ismaleDisabled"/>
                    男
                  </label>
                  <label>
                    <input type="radio" name="gender" value="female" class="radio radio-primary"
                           v-model="selectedGender"
                           :disabled="isFemaleDisabled"/>
                    女
                  </label>
                </div>


              </div>
              <div class="space-y-2"><label
                  class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                  style="margin-bottom:20px" for="message">用户头像</label>
                <div class="flex items-center"> <!-- Flexbox 容器 -->
                  <!-- 第一个按钮 -->
                  <div class="avatar">
                    <div class="w-24 rounded-full" style="width: 50px;height: 50px">
                      <img :src="startAvatar"/>
                    </div>
                    <!--                    上传新的用户头像-->
                    <input type="file" class="file-input file-input-bordered file-input-secondary w-full max-w-xs"
                           @change="handleFileUpload" style="margin-top: 2%;margin-left: 50px;width: 60%"/>
                  </div>
                </div>
              </div>
              <button
                  class="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 hover:bg-primary/90 h-10 px-4 py-2 bg-gray-800 text-white"
                  type="submit" @click="updatePass">提交修改
              </button>
              <button class="btn" style="margin-left: 50%;" @click="cancel">取消修改</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- 拟态框 -->
  <div class="modal" role="dialog" id="my_modal_8">
    <div class="bg-white p-6 rounded-lg shadow-md max-w-sm" style="padding-left: 5%;padding-right:5%">
      <div class="flex items-center space-x-3 mb-6"><img src="@/assets/police.png" alt="App icon" class="h-10 w-10"
                                                         width="10" height="10"
                                                         style="aspect-ratio: 100 / 100; object-fit: cover;">
        <h3 class="text-lg font-semibold">实名认证</h3>
      </div>
      <label class="form-control w-full max-w-xs" style="margin-bottom: 8%;">
        <div class="label">
          <span class="label-text" style="margin-left: -120%;">您的姓名?</span>
        </div>
        <input type="text" placeholder="真实姓名" class="input input-bordered w-full max-w-xs" v-model="name"/>
        <div class="label">
          <span class="label-text" style="margin-left: -90%;">您的身份证?</span>
        </div>
        <input type="text" placeholder="身份证" class="input input-bordered w-full max-w-xs" v-model="id_card"/>
      </label>
      <button @click="accountBindUserInfo"
              class="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 h-10 px-4 py-2 bg-blue-500 hover:bg-blue-600 text-white mb-4">
        实名认证
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
             stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="ml-2">
          <path d="M5 12h14"></path>
          <path d="m12 5 7 7-7 7"></path>
        </svg>
      </button>
      <button class="btn" width="24" height="24" style="width: 85px;height:10px;margin-left:20px"
              @click="cancel">取消
      </button>
      <p class="text-xs text-gray-500" style="text-align:center">区块链养老保险平台</p>
    </div>
  </div>

  <div class="success" style="position: absolute; right:7%;top:3%;z-index:100" v-show="alert">
    <div role="alert" class="alert alert-error">
      <svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
              d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"/>
      </svg>
      <span v-text="message"></span>
    </div>
  </div>
  <div class="success" style="position: absolute; right:5%;top:3%;z-index:100" v-show="success">
    <div role="alert" class="alert alert-success">
      <svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
              d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
      </svg>
      <span v-text="message"></span>
    </div>
  </div>

  <router-view v-slot="{ Component }">
    <transition name="fade">
      <component :is="Component"/>
    </transition>
  </router-view>
</template>

<script>
export default {
  data() {
    return {
      name: '',
      id_card: '',
      message: '',
      alert: false,
      success: false,
      selectedGender: 'male', // 默认选中男性
      isFemaleDisabled: false,//
      ismaleDisabled: false,
      colorStart: '',
      colorEnd: '',
      avatar: '',
      username: '',
      age: 0,
      gender: '',
      password: '',
      startAvatar:'',
    }
  },
  created() {
    this.startAvatar=localStorage.getItem('avatar')
    this.colorStart = localStorage.getItem("colorStart")
    this.colorEnd = localStorage.getItem("colorEnd")
    // 获取用户信息
    this.$axios.get("/user/getUserInfo").then((rep) => {
      console.log(rep);
      this.message = rep.data.message
      if (rep.data.code != 200) {
        return
      }
      this.username = rep.data.userInfo.username
      this.age = rep.data.userInfo.age
      this.gender = rep.data.userInfo.gender

      if (this.gender == '男') {
        this.isFemaleDisabled = true
      } else {
        this.selectedGender = "female"
        this.ismaleDisabled = true
      }
    }).catch((rep) => {
      console.log(rep);
    })
  },
  methods: {
    updatePass() {
      if (this.password.length == 0&&this.avatar.length==0) {
        window.location.href = '#';
        // 将生成的颜色存入 localStorage
        localStorage.setItem("colorStart", this.colorStart);
        localStorage.setItem("colorEnd", this.colorEnd);
        this.message = "修改用户信息成功"
        this.success = true
        setTimeout(() => {
          this.success = false;
        }, 1000); // 1秒后执行
      } else {
        window.location.href = '#';
        // 将生成的颜色存入 localStorage
        localStorage.setItem("colorStart", this.colorStart);
        localStorage.setItem("colorEnd", this.colorEnd);
        this.$axios.post("/user/updatePassword", {
          "password": this.password,
          "avatar": this.avatar,
        }).then((rep) => {
          this.message = rep.data.message
          if (rep.data.code != 200) {
            this.alert = true
            setTimeout(() => {
              this.alert = false;
            }, 1000); // 1秒后执行
            return
          }
          // 更新本地头像缓存
          localStorage.setItem("avatar",this.avatar);
          this.success = true
          this.password = ''
          setTimeout(() => {
            this.success = false;
          }, 1000); // 1秒后执行
        }).catch((rep) => {
          console.log(rep);
        })
      }
    },
    // 实名认证
    accountBindUserInfo() {
      if (this.name.length == 0 || this.id_card.length != 18) {
        window.location.href = '#';
        this.message = "填写信息有误"
        this.alert = true
        this.name = ''
        this.id_card = ''
        setTimeout(() => {
          this.alert = false;
        }, 1000); // 1秒后执行
        return
      }
      this.$axios.post("/user/UserRealNameHandler", {
        name: this.name,
        idCard: this.id_card
      }).then((rep) => {
        window.location.href = '#';
        this.message = rep.data.message
        this.name = ''
        this.id_card = ''
        if (rep.data.code != 200) {
          this.alert = true
          setTimeout(() => {
            this.alert = false;
          }, 1000); // 1秒后执行
          return
        }
        localStorage.setItem("user_auth", "true")
        window.location.reload();
        this.success = true
        setTimeout(() => {
          this.success = false;
        }, 1000); // 1秒后执行
      }).catch((rep) => {
        console.log(rep);
      })
    },
    cancel() {
      window.location.href = '#';
    },
    // 刷新颜色
    refreshColor() {
      this.colorStart = this.getRandomColor();
      this.colorEnd = this.getRandomColor();
    },
    getRandomColor() {
      const letters = '0123456789ABCDEF';
      let color = '#';
      for (let i = 0; i < 6; i++) {
        color += letters[Math.floor(Math.random() * 16)];
      }
      return color;
    },
    getAvatarStyle(user) {
      // 尝试从 localStorage 中获取颜色
      let colorStart = this.colorStart;
      let colorEnd = this.colorEnd;

      // 如果没有找到缓存的颜色，生成新的颜色
      if (!colorStart || !colorEnd) {
        colorStart = this.getRandomColor();
        colorEnd = this.getRandomColor();

        // 将生成的颜色存入 localStorage
        localStorage.setItem("colorStart", colorStart);
        localStorage.setItem("colorEnd", colorEnd);
      }

      // 通过指定色标的位置来增加渐变尺度
      // 比如从 10% 的位置开始到 90% 的位置结束
      return {
        background: `linear-gradient(to right, ${colorStart}, ${colorEnd})`,
        width: '52px',
        height: '52px',
        borderRadius: '50%',
        aspectRatio: '1 / 1',
        objectFit: 'cover',
      };
    },
    handleFileUpload(event) {
      const file = event.target.files[0]; // 获取用户选择的文件
      if (file) {
        const reader = new FileReader();

        reader.onload = (e) => {
          // 设置头像
          this.avatar = e.target.result;
          console.log(e.target.result); // 这里是文件的Base64编码
          // 你可以在这里处理转换后的Base64字符串，比如将其发送到服务器
        };

        reader.readAsDataURL(file); // 将文件读取为Data URL
      }
    },
  }
}

</script>

<style scoped>
/* 渐变设置 */
.fade-enter-from,
.fade-leave-to {
  transform: translateX(20px);
  opacity: 0;
}

.fade-enter-to,
.fade-leave-from {
  opacity: 1;
}

.fade-enter-active {
  transition: all 0.7s ease;
}

.fade-leave-active {
  transition: all 0.3s cubic-bezier(1, 0.6, 0.6, 1);
}

.radio-group {
  display: flex;
  /* 设置为 flex 布局 */
  flex-wrap: nowrap;
  /* 防止子元素换行 */
}

label {
  display: flex;
  align-items: center;
  margin-right: 10px;
  white-space: nowrap;
}

.radio {
  margin-right: 4px;
}
</style>
