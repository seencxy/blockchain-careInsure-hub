<template>
  <div class="app">
    <!-- Put this part before </body> tag -->
    <div class="modal" role="dialog" id="create">
      <div class="modal-box">
        <h3 class="font-bold text-lg mb-4">添加套餐</h3>
        <span>套餐名称： </span> <input type="text" placeholder="输入套餐名" v-model="combo.Name"
                                       class="input input-bordered w-full mb-2 max-w-xs"/>
        <br>
        <span>套餐价格： </span><input type="number" placeholder="输入套餐价格(人民币)" v-model.number="combo.Price"
                                      class="input input-bordered w-full mb-2 max-w-xs"/><br>
        <span>套餐开始年龄： </span><input type="number" placeholder="输入开始年龄(岁)" v-model.number="combo.StartYear"
                                          class="input input-bordered w-full mb-2 max-w-xs"/><br>
        <span>套餐结束年龄： </span> <input type="number" placeholder="输入结束年龄(岁)" v-model.number="combo.EndYear"
                                           class="input input-bordered w-full mb-2 max-w-xs"/><br>
        <span>每月回退养老金： </span><input type="number" placeholder="输入每月反馈费用(人民币)"
                                            v-model.number="combo.MonthFee"
                                            class="input input-bordered w-full mb-2 max-w-xs"/><br>
        <label class="flex items-center space-x-2 mb-2">
          <span>高额医疗覆盖：</span> <input type="checkbox" v-model="combo.HighMedicalCoverage"
                                            class="toggle toggle-accent"/>
        </label>
        <span>支持退款时间： </span><input type="number" placeholder="退款时间(月)" v-model.number="combo.RefundPeriod"
                                          class="input input-bordered w-full mb-4 max-w-xs"/><br>

        <span>商品简介： </span>
        <div v-for="(desc, index) in combo.Description" :key="index" class="mb-2">
          <input type="text" v-model="desc.description" placeholder="输入简介"
                 class="input input-bordered w-full max-w-xs"/>
          <button @click="removeDescription(index)" class="btn btn-error btn-xs ml-2">删除</button>
        </div>
        <button @click="addDescription" class="btn btn-primary btn-sm">添加简介</button>

        <div class="modal-action">
          <button @click="cancel" class="btn">取消</button>
          <button @click="submitCombo" class="btn">确认</button>
        </div>
      </div>
    </div>
    <div class="flex">
      <aside class="h-screen w-56 bg-gray-100 text-gray-800 p-4">
        <div class="flex items-center mb-4 space-x-1">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"
               fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
               stroke-linejoin="round" class="h-6 w-6">
            <path d="M3 9h18v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V9Z"></path>
            <path d="m3 9 2.45-4.9A2 2 0 0 1 7.24 3h9.52a2 2 0 0 1 1.8 1.1L21 9"></path>
            <path d="M12 3v6"></path>
          </svg>
          <h1 class="text-lg font-medium">管理后台</h1>
        </div>
        <nav class="space-y-2">
          <button
              class="w-full flex items-center space-x-2 hover:bg-gray-200 active:bg-gray-300 py-2 px-2 rounded-lg text-gray-500"
              @click="navigateToPackage('/admin/home')">
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
                class="w-4 h-4"
            >
              <path d="m3 9 9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"></path>
              <polyline points="9 22 9 12 15 12 15 22"></polyline>
            </svg>
            <span class="text-sm font-medium">用户管理</span>
          </button>
          <button
              class="w-full flex items-center space-x-2 hover:bg-gray-200 active:bg-gray-300 py-2 px-2 rounded-lg text-gray-500"
              @click="navigateToPackage('/admin/pay')">
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
                class="w-4 h-4"
            >
              <path d="M21 12V7H5a2 2 0 0 1 0-4h14v4"></path>
              <path d="M3 5v14a2 2 0 0 0 2 2h16v-5"></path>
              <path d="M18 12a2 2 0 0 0 0 4h4v-4Z"></path>
            </svg>
            <span class="text-sm font-medium">支付管理</span>
          </button>
          <button
              class="w-full flex items-center space-x-2 bg-gray-200 active:bg-gray-300 py-2 px-2 rounded-lg text-gray-800"
              @click="navigateToPackage('/admin/package')">
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
                class="w-4 h-4"
            >
              <path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"></path>
              <circle cx="9" cy="7" r="4"></circle>
              <path d="M22 21v-2a4 4 0 0 0-3-3.87"></path>
              <path d="M16 3.13a4 4 0 0 1 0 7.75"></path>
            </svg>
            <span class="text-sm font-medium">套餐管理</span>
          </button>
          <router-link to="/login">
            <button
                style="display: flex; justify-content: center; align-items: center; width: 100%; background-color: #007bff; color: white; padding: 8px; border-radius: 8px; margin-top: 65vh;"
            >
              <span style="text-align: center; font-size: small;font-weight: bold">退出登录</span>
            </button>
          </router-link>
        </nav>
      </aside>
      <main class="flex-grow p-6">
        <div class="flex justify-between items-center mb-4">
          <h1 class="text-lg font-medium">套餐管理</h1>
          <button
              class="justify-center whitespace-nowrap text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 hover:bg-primary/90 h-10 px-4 py-2 bg-gray-800 text-white rounded-lg flex items-center space-x-2"
              type="button" @click="openCreate()"
          >
            <span>添加套餐</span>
          </button>
        </div>
        <div class="overflow-auto" style="height: 85vh">
          <table class="table table-xs table-pin-rows table-pin-cols">
            <!-- head -->
            <thead>
            <tr>
              <th>套餐ID</th>
              <th>套餐名称</th>
              <th>套餐价格</th>
              <th>套餐购买数量</th>
              <th>套餐开始年龄</th>
              <th>套餐结束年龄</th>
              <th>月反馈养老金</th>
              <th>是否支持高医疗</th>
              <th>套餐状态</th>
              <th>操作</th>
            </tr>
            </thead>
            <tbody>
            <!-- row 1 -->
            <tr v-for="packageInfo in packages" :key="packageInfo.ID" style="height: 60px">
              <td>{{ packageInfo.ID }}</td>
              <td>
                {{ packageInfo.name }}
              </td>
              <td>{{ packageInfo.price }}¥</td>

              <td>{{ packageInfo.start_year }}</td>
              <td>{{ packageInfo.end_year }}</td>
              <td>{{ packageInfo.end_year }}</td>
              <td>{{ packageInfo.month_fee }}</td>
              <td>{{ packageInfo.highMedicalCoverage }}</td>
              <td>{{ showStatus(packageInfo.description[0].ComboID) }}</td>
              <td>
                <button class="btn btn-ghost btn-xs" v-show="packageInfo.description[0].ComboID===1"
                        @click="updateComboStatus(packageInfo.ID)">停用套餐
                </button>
                <button class="btn btn-ghost btn-xs" v-show="packageInfo.description[0].ComboID===0"
                        @click="updateComboStatus(packageInfo.ID)">启用套餐
                </button>
              </td>
            </tr>
            </tbody>
            <tfoot>
            <tr>
              <th>套餐ID</th>
              <th>套餐名称</th>
              <th>套餐价格</th>
              <th>套餐购买数量</th>
              <th>套餐开始年龄</th>
              <th>套餐结束年龄</th>
              <th>月反馈养老金</th>
              <th>是否支持高医疗</th>
              <th>套餐状态</th>
              <th>操作</th>
            </tr>
            </tfoot>

          </table>
        </div>
      </main>
    </div>
  </div>
</template>

<script>
export default {
  name: 'package',
  created() {
    this.getAllCombo();
  },
  methods: {
    navigateToPackage(router) {
      this.$router.push(router);
    },
    openCreate() {
      window.location.href = '#create'
    },
    addDescription() {
      this.combo.Description.push({text: ''});
    },
    removeDescription(index) {
      this.combo.Description.splice(index, 1);

    },
    getAllCombo() {
      this.$axios.get("/main/getAllCombo").then((res) => {
        if (res.data.code === 200) {
          this.packages = res.data.data
        } else {
          this.$message.error("获取订单信息失败: " + res.data.message)
        }
        console.log(res)
      }).catch((res) => {
        console.log(res)
      })
    }
    ,
    submitCombo() {
      // 在这里处理表单提交逻辑，例如通过API发送数据
      console.log('提交的套餐数据:', this.combo);
      this.$axios.post("/main/createCombo", {
        name: this.combo.Name,
        price: this.combo.Price,
        start_year: this.combo.StartYear,
        end_year: this.combo.EndYear,
        month_fee: this.combo.MonthFee,
        highMedicalCoverage: this.combo.HighMedicalCoverage,
        refundPeriod: this.combo.RefundPeriod,
        description: this.combo.Description
      }).then((res) => {
        if (res.data.code === 200) {
          this.$message.success("创建套餐成功: " + res.data.transactionHash)
        } else {
          this.$message.error("创建套餐失败")
        }
        window.location.href = '#'
        console.log(res)
      }).catch((res) => {
        console.log(res)
      })
    },
    cancel() {
      window.location.href = '#'
    },
    showStatus(data) {
      if (data === 1) {
        return "使用中"
      } else {
        return "停用"
      }
    },
    updateComboStatus(id) {
      this.$axios.post("/main/updateComboStatus", {
        ID: id
      }).then((res) => {
        if (res.data.code === 200) {
          this.$message.success("修改状态成功")
        } else {
          this.$message.error("修改状态失败")
        }
        console.log(res)
      }).catch((res) => {
        console.log(res)
      })
    }
  },
  data() {
    return {
      packages: [],
      combo: {
        Name: '',
        Description: [{text: ''}],
        Price: null,
        StartYear: null,
        EndYear: null,
        MonthFee: null,
        HighMedicalCoverage: false,
        RefundPeriod: null,
      },
    }
  }
}
</script>

<style scoped>

</style>