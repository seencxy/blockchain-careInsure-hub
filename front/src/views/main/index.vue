<template>
  <!-- Put this part before </body> tag -->
  <div class="modal" role="dialog" id="question">
    <div class="modal-box">
      <h3 class="font-bold text-lg" v-text="question"></h3>
      <p class="py-4" v-text="answer"></p>
      <div class="modal-action">
        <a href="#" class="btn">好的</a>
      </div>
    </div>
  </div>
  <div class="app">
    <div class="grid h-screen min-h-screen w-full overflow-hidden lg:grid-cols-[280px_1fr]">
      <div class="hidden border-r bg-gray-100/40 lg:block dark:bg-gray-800/40">
        <div class="flex flex-col gap-2">
          <div class="flex h-[60px] items-center px-6"><a class="flex items-center gap-2 font-semibold"
                                                          href="#">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"
                 fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                 stroke-linejoin="round" class="h-6 w-6">
              <path d="M3 9h18v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V9Z"></path>
              <path d="m3 9 2.45-4.9A2 2 0 0 1 7.24 3h9.52a2 2 0 0 1 1.8 1.1L21 9"></path>
              <path d="M12 3v6"></path>
            </svg>
            <span class="">区块链养老保险平台</span></a></div>
          <div class="flex-1">
            <nav class="grid items-start px-4 text-sm font-medium">
              <router-link to="/"><a
                  class="flex items-center gap-3 rounded-lg bg-gray-100 px-3 py-2 text-gray-900  transition-all hover:text-gray-900 dark:bg-gray-800 dark:text-gray-50 dark:hover:text-gray-50"
                  href="#">
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24"
                     viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                     stroke-linecap="round" stroke-linejoin="round" class="h-4 w-4">
                  <path d="m3 9 9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"></path>
                  <polyline points="9 22 9 12 15 12 15 22"></polyline>
                </svg>
                主页
              </a></router-link>
              <router-link to="order"><a
                  class="flex items-center gap-3 rounded-lg px-3 py-2 text-gray-500 transition-all hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-50"
                  href="#">
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24"
                     viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                     stroke-linecap="round" stroke-linejoin="round" class="h-4 w-4">
                  <circle cx="8" cy="21" r="1"></circle>
                  <circle cx="19" cy="21" r="1"></circle>
                  <path
                      d="M2.05 2.05h2l2.66 12.42a2 2 0 0 0 2 1.58h9.78a2 2 0 0 0 1.95-1.57l1.65-7.43H5.12">
                  </path>
                </svg>
                交易状况
              </a></router-link>
              <router-link to="product"><a
                  class="flex items-center gap-3 rounded-lg px-3 py-2 text-gray-500 transition-all hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-50"
                  href="#">
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24"
                     viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                     stroke-linecap="round" stroke-linejoin="round" class="h-4 w-4">
                  <path d="m7.5 4.27 9 5.15"></path>
                  <path
                      d="M21 8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16Z">
                  </path>
                  <path d="m3.3 7 8.7 5 8.7-5"></path>
                  <path d="M12 22V12"></path>
                </svg>
                产品
              </a></router-link>
              <router-link to="customer"><a
                  class="flex items-center gap-3 rounded-lg px-3 py-2 text-gray-500 transition-all hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-50"
                  href="#">
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24"
                     viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                     stroke-linecap="round" stroke-linejoin="round" class="h-4 w-4">
                  <path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"></path>
                  <circle cx="9" cy="7" r="4"></circle>
                  <path d="M22 21v-2a4 4 0 0 0-3-3.87"></path>
                  <path d="M16 3.13a4 4 0 0 1 0 7.75"></path>
                </svg>
                顾客
              </a></router-link>
              <router-link to="analytic"><a
                  class="flex items-center gap-3 rounded-lg px-3 py-2 text-gray-500 transition-all hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-50"
                  href="#">
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24"
                     viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                     stroke-linecap="round" stroke-linejoin="round" class="h-4 w-4">
                  <path d="M3 3v18h18"></path>
                  <path d="m19 9-5 5-4-4-3 3"></path>
                </svg>
                分析
              </a></router-link>
              <router-link to="nodeInfo"><a
                  class="flex items-center gap-3 rounded-lg px-3 py-2 text-gray-500 transition-all hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-50"
                  href="#">
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="24"
                    height="24"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    class="h-4 w-4"
                >
                  <path d="m7.5 4.27 9 5.15"></path>
                  <path
                      d="M21 8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16Z"></path>
                  <path d="m3.3 7 8.7 5 8.7-5"></path>
                  <path d="M12 22V12"></path>
                </svg>
                节点信息
              </a></router-link>
            </nav>
          </div>
        </div>
      </div>
      <div class="flex flex-col">
        <header class="flex h-14 lg:h-[60px] items-center gap-4 border-b bg-gray-100/40 px-6 dark:bg-gray-800/40"><a
            class="lg:hidden" href="#">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24"
               viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
               stroke-linejoin="round" class="h-6 w-6">
            <path d="M3 9h18v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V9Z"></path>
            <path d="m3 9 2.45-4.9A2 2 0 0 1 7.24 3h9.52a2 2 0 0 1 1.8 1.1L21 9"></path>
            <path d="M12 3v6"></path>
          </svg>
          <span class="sr-only">Home</span></a>
          <div class="flex-1">
            <h1 class="font-semibold text-lg">主页</h1>
          </div>
          <div class="flex flex-1 items-center gap-4 md:ml-auto md:gap-2 lg:gap-4">
            <form class="ml-auto flex-1 sm:flex-initial">
            </form>
            <div class="dropdown dropdown-end">
              <div tabindex="0" role="button" class="btn btn-ghost btn-circle avatar">
                <div class="w-12 rounded-full">
                  <img alt="Tailwind CSS Navbar component" :src="avatar" style="width: 100%"/>
                </div>
              </div>
              <ul tabindex="0"
                  class="mt-3 z-[1] p-2 shadow menu menu-sm dropdown-content bg-base-100 rounded-box w-52">
                <li @click="openVerified">
                  <a class="justify-between">
                    实名认证
                    <span class="badge" v-if="isUserAuthenticated">已认证</span>
                    <span class="badge" v-else>未认证</span>
                  </a>
                </li>
                <li @click="openUpdate"><a>修改用户信息</a></li>
                <li @click="logout"><a>登出</a></li>
                <li @click="userLogout"><a>账号注销</a></li>
              </ul>
            </div>
          </div>
        </header>
        <main class="flex flex-1 flex-col gap-4 p-4 md:gap-8 md:p-6">

          <div class="flex justify-center items-center min-h-screen" style="margin-top:-10%">
            <div class="rounded-lg border bg-card text-card-foreground shadow-lg w-[550px]" data-v0-t="card">
              <div class="p-6">
                <h2 class="text-2xl font-semibold">
                  区块链养老保险平台相关问题
                </h2>
                <div class="w-full mt-4" collapsible="" data-orientation="vertical">
                  <div data-state="closed" data-orientation="vertical" class="border-b"
                       @click="getAnswer(1)">
                    <h3 data-orientation="vertical" data-state="closed" class="flex">
                      <button
                          type="button" aria-controls="radix-:R2dunnnla:" aria-expanded="false"
                          data-state="closed" data-orientation="vertical" id="radix-:Rdunnnla:"
                          class="flex flex-1 items-center justify-between py-4 font-medium transition-all hover:underline [&amp;[data-state=open]>svg]:rotate-180 hover:underline-none"
                          data-radix-collection-item="">
                        这个区块链养老保险平台的主要目标是什么？
                        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24"
                             viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                             stroke-linecap="round" stroke-linejoin="round"
                             class="h-4 w-4 shrink-0 transition-transform duration-200">
                          <path d="m6 9 6 6 6-6"></path>
                        </svg>
                      </button>
                    </h3>
                  </div>
                  <div data-state="closed" data-orientation="vertical" class="border-b" @click="getAnswer(2)">
                    <h3 data-orientation="vertical" data-state="closed" class="flex">
                      <button
                          type="button" aria-controls="radix-:R2lunnnla:" aria-expanded="false"
                          data-state="closed" data-orientation="vertical" id="radix-:Rlunnnla:"
                          class="flex flex-1 items-center justify-between py-4 font-medium transition-all hover:underline [&amp;[data-state=open]>svg]:rotate-180 hover:underline-none"
                          data-radix-collection-item="">
                        平台将如何利用区块链技术改善养老服务？
                        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24"
                             viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                             stroke-linecap="round" stroke-linejoin="round"
                             class="h-4 w-4 shrink-0 transition-transform duration-200">
                          <path d="m6 9 6 6 6-6"></path>
                        </svg>
                      </button>
                    </h3>
                    <div data-state="closed" id="radix-:R2lunnnla:" hidden="" role="region"
                         aria-labelledby="radix-:Rlunnnla:" data-orientation="vertical"
                         class="overflow-hidden text-sm transition-all data-[state=closed]:animate-accordion-up data-[state=open]:animate-accordion-down"
                         style="--radix-accordion-content-height: var(--radix-collapsible-content-height); --radix-accordion-content-width: var(--radix-collapsible-content-width);">
                    </div>
                  </div>
                  <div data-state="closed" data-orientation="vertical" class="border-b" @click="getAnswer(3)">
                    <h3 data-orientation="vertical" data-state="closed" class="flex">
                      <button
                          type="button" aria-controls="radix-:R2tunnnla:" aria-expanded="false"
                          data-state="closed" data-orientation="vertical" id="radix-:Rtunnnla:"
                          class="flex flex-1 items-center justify-between py-4 font-medium transition-all hover:underline [&amp;[data-state=open]>svg]:rotate-180 hover:underline-none"
                          data-radix-collection-item="">
                        相比于传统养老保险服务，平台的主要竞争优势是什么？
                        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24"
                             viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                             stroke-linecap="round" stroke-linejoin="round"
                             class="h-4 w-4 shrink-0 transition-transform duration-200">
                          <path d="m6 9 6 6 6-6"></path>
                        </svg>
                      </button>
                    </h3>
                    <div data-state="closed" id="radix-:R2tunnnla:" hidden="" role="region"
                         aria-labelledby="radix-:Rtunnnla:" data-orientation="vertical"
                         class="overflow-hidden text-sm transition-all data-[state=closed]:animate-accordion-up data-[state=open]:animate-accordion-down"
                         style="--radix-accordion-content-height: var(--radix-collapsible-content-height); --radix-accordion-content-width: var(--radix-collapsible-content-width);">
                    </div>
                  </div>
                  <div data-state="closed" data-orientation="vertical" class="border-b" @click="getAnswer(4)">
                    <h3 data-orientation="vertical" data-state="closed" class="flex">
                      <button
                          type="button" aria-controls="radix-:R35unnnla:" aria-expanded="false"
                          data-state="closed" data-orientation="vertical" id="radix-:R15unnnla:"
                          class="flex flex-1 items-center justify-between py-4 font-medium transition-all hover:underline [&amp;[data-state=open]>svg]:rotate-180 hover:underline-none"
                          data-radix-collection-item="">
                        平台如何确保用户数据的安全和隐私？
                        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24"
                             viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                             stroke-linecap="round" stroke-linejoin="round"
                             class="h-4 w-4 shrink-0 transition-transform duration-200">
                          <path d="m6 9 6 6 6-6"></path>
                        </svg>
                      </button>
                    </h3>
                    <div data-state="closed" id="radix-:R35unnnla:" hidden="" role="region"
                         aria-labelledby="radix-:R15unnnla:" data-orientation="vertical"
                         class="overflow-hidden text-sm transition-all data-[state=closed]:animate-accordion-up data-[state=open]:animate-accordion-down"
                         style="--radix-accordion-content-height: var(--radix-collapsible-content-height); --radix-accordion-content-width: var(--radix-collapsible-content-width);">
                    </div>
                  </div>
                  <div data-state="closed" data-orientation="vertical" class="border-b-0" @click="getAnswer(5)">
                    <h3 data-orientation="vertical" data-state="closed" class="flex">
                      <button
                          type="button" aria-controls="radix-:R3dunnnla:" aria-expanded="false"
                          data-state="closed" data-orientation="vertical" id="radix-:R1dunnnla:"
                          class="flex flex-1 items-center justify-between py-4 font-medium transition-all hover:underline [&amp;[data-state=open]>svg]:rotate-180 hover:underline-none"
                          data-radix-collection-item="">
                        平台如何利用区块链来确保交易的透明度和安全性？
                        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24"
                             viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                             stroke-linecap="round" stroke-linejoin="round"
                             class="h-4 w-4 shrink-0 transition-transform duration-200">
                          <path d="m6 9 6 6 6-6"></path>
                        </svg>
                      </button>
                    </h3>
                    <div data-state="closed" id="radix-:R3dunnnla:" hidden="" role="region"
                         aria-labelledby="radix-:R1dunnnla:" data-orientation="vertical"
                         class="overflow-hidden text-sm transition-all data-[state=closed]:animate-accordion-up data-[state=open]:animate-accordion-down"
                         style="--radix-accordion-content-height: var(--radix-collapsible-content-height); --radix-accordion-content-width: var(--radix-collapsible-content-width);">
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </main>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "index",
  data() {
    return {
      isUserAuthenticated: false,
      question: '',
      answer: '',
      avatar: '',
      faqs: {
        1: {
          question: "这个区块链养老平台的主要目标是什么？",
          answer: "这个区块链养老平台的主要目标是利用区块链技术提供一个更安全、透明、高效的养老资金管理和服务平台。目标包括简化资金流程、减少欺诈风险、提高资金使用的透明度，以及为用户提供更个性化的养老规划服务。"
        },
        2: {
          question: "平台将如何利用区块链技术改善养老服务？",
          answer: "平台将通过智能合约、透明性、安全性、去中心化和数据整合等方式利用区块链技术来改善养老服务。"
        },
        3: {
          question: "3. 相比于传统养老服务，平台的主要竞争优势是什么？",
          answer: `与传统养老服务相比，该平台的主要竞争优势包括：
增强的安全性和透明度：通过区块链技术确保交易记录的不可篡改性和可追溯性。
降低运营成本：自动化的智能合约减少人工干预，降低管理费用。
用户自主权：用户对自己的资金和数据拥有更大的控制权。
灵活性和可定制化：提供更多样化的投资选项和个性化的养老规划。`
        },
        4: {
          question: "4. 平台如何确保用户数据的安全和隐私？",
          answer: `平台通过以下方式确保用户数据的安全和隐私：

加密技术：对敏感数据进行加密，确保数据在传输和存储过程中的安全。
访问控制：严格的权限管理和访问控制，确保只有授权用户可以访问数据。
隐私保护协议：运用区块链的隐私保护技术，如零知识证明，保护用户隐私。
数据分割存储：将数据分割成多部分，分散存储，减少数据泄露的风险。`
        },
        5: {
          question: "5. 平台如何利用区块链来确保交易的透明度和安全性？",
          answer: `平台通过以下方式利用区块链确保交易的透明度和安全性：

不可篡改的记录：所有交易都在区块链上永久记录，无法被篡改，确保交易历史的真实性。
公开透明的过程：交易过程和结果对所有相关方公开，提高了交易的透明度。
去中心化验证：通过网络中的多个节点进行交易验证，减少欺诈和错误。
安全协议：采用高级的加密算法和安全协议保护交易和数据不受攻击。`
        }
      },

    }
  },
  created() {
    const authValue = localStorage.getItem('user_auth');
    this.isUserAuthenticated = authValue === 'true'; // 确保将字符串 'true' 转换为布尔值 true
    // 获取用户头像
    this.avatar = localStorage.getItem("avatar");
  },
  methods: {
    // 账号注销接口
    userLogout() {
      this.$axios.post('/user/logout', {}, {
        headers: {
          'Authorization': `Bearer ` + localStorage.getItem('auth_token')
        }
      }).then((res) => {
        console.log(res)
      }).catch((error) => {
        console.error(error)
      })

      // 删除所有缓存
      localStorage.removeItem("auth_token")
      localStorage.removeItem("user_auth")
      // 将路由切换到 /login
      this.$router.push('/login')
    },
    getAnswer(i) {
      if (this.faqs[i]) {
        this.question = this.faqs[i].question;
        this.answer = this.faqs[i].answer;
      } else {
        this.question = '问题编号不正确。';
        this.answer = '';
      }
      window.location.href = '#question';
    },
    // 账号登出
    logout() {
      // 删除所有缓存
      localStorage.removeItem("auth_token")
      localStorage.removeItem("user_auth")
      // 将路由切换到 /login
      this.$router.push('/login')
    },
    openVerified() {
      if (this.isUserAuthenticated == false) {
        window.location.href = '#my_modal_8';
      }
    },
    openUpdate() {
      if (this.isUserAuthenticated == true) {
        window.location.href = '#updateUserInfo';
      }
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
      let colorStart = localStorage.getItem("colorStart");
      let colorEnd = localStorage.getItem("colorEnd");

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
        width: '32px',
        height: '32px',
        borderRadius: '50%',
        aspectRatio: '1 / 1',
        objectFit: 'cover'
      };
    },
  }

};
</script>

<style scoped>
.avatar {
  display: inline-block;
  overflow: hidden;
  /* 如果你要在里面放图标或者文字，建议设置 overflow */
  /* 其他样式 */
}
</style>