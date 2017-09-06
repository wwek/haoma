<template>
  <div class="main" id="phonsreach">
    <el-row>
      <el-col :span="24"><div class="grid-content">
        <h1 class="title">电话号码标记查询</h1>
      </div></el-col>
    </el-row>
    <el-row>
      <el-col :span="4" class="plist">
          <el-form label-width="100px">
            <el-button type="primary" v-on:click="dosearch" class="search" title="点击批量查询后右侧看结果">批量查询</el-button>
            <textarea  v-model="phoneNumberList" placeholder="输入电话号码列表,支持手机和固话,一行一个" type="textarea" rows="2" autocomplete="off" validateevent="true" class="text"></textarea>
          </el-form>
      </el-col>
      <el-col :span="20">
        <div class="grid-content bg-purple-light">
          <el-table
            :data="tableData"
            stripe
            sortable
            style="width: 100%">
            <el-table-column
              prop="phone_number"
              label="号码"
              width="130">
            </el-table-column>
            <el-table-column
              prop="from1"
              label="渠道1"
              width="">
            </el-table-column>
            <el-table-column
              prop="location1"
              label="1-归属地"
              width=""
              >
            </el-table-column>
            <el-table-column
              prop="tag_name1"
              label="1-标记"
              width=""
            >
            </el-table-column>
            <el-table-column
              prop="tag_cnt1"
              label="1-次"
              width="80"
            >
            </el-table-column>
            <el-table-column
              prop="from2"
              label="渠道2"
              width="">
            </el-table-column>
            <el-table-column
              prop="location2"
              label="2-归属地"
              width="">
            </el-table-column>
            <el-table-column
              prop="tag_name2"
              label="2-标记"
              width=""
            >
            </el-table-column>
            <el-table-column
              prop="tag_cnt2"
              label="2-次"
              width="80"
            >
            </el-table-column>
            <el-table-column
              prop="from3"
              label="渠道3"
              width="">
            </el-table-column>
            <el-table-column
              prop="location3"
              label="3-归属地"
              width="">
            </el-table-column>
            <el-table-column
              prop="tag_name3"
              label="3-标记"
              width=""
            >
            </el-table-column>
            <el-table-column
              prop="tag_cnt3"
              label="3-次"
              width="80"
            >
            </el-table-column>
          </el-table>

        </div>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="24"><div class="grid-content">
        <div class="footer">
        <span class="cp">&copy; Wwek </span> <a target="_blank" href="https://github.com/wwek/haoma">Haoma 电话号码标记查询</a>
        </div>
      </div></el-col>
    </el-row>
  </div>
</template>

<script>
  export default {
    name: 'haoma',
    data () {
      return {
        phoneNumberList: '053266114000',
        tableData: []
      }
    },
    methods: {
      dosearch: function () {
        this.tableData = []
        var phoneArr = this.phoneNumberList.split('\n')
        for (var i = 0; i < phoneArr.length; i++) {
          var pn = this.$_.trim(phoneArr[i])
          var that = this
          this.$http.get('/v1/phone/tag', {
            params: {
              phonenumber: pn
            }
          })
            .then(function (response) {
              var oneData = that.apiToRow(response.data.data)
              console.log(that.tableData)
              console.log(oneData)
              that.tableData.push(oneData)
            })
            .catch(function (error) {
              console.log(error)
            })
        }
      },
      apiToRow (rdata) {
        var oneData = {}
        oneData.phone_number = rdata[0].phone_number
        oneData.from1 = rdata[0].from
        oneData.location1 = rdata[0].location.province + ' ' + rdata[0].location.city
        oneData.tag_name1 = rdata[0].tag.tag_name
        oneData.tag_cnt1 = rdata[0].tag.tag_cnt
        oneData.from2 = rdata[1].from
        oneData.location2 = rdata[1].location.province + ' ' + rdata[1].location.city
        oneData.tag_name2 = rdata[1].tag.tag_name
        oneData.tag_cnt2 = rdata[1].tag.tag_cnt
        oneData.from3 = rdata[2].from
        oneData.location3 = rdata[2].location.province + ' ' + rdata[2].location.city
        oneData.tag_name3 = rdata[2].tag.tag_name
        oneData.tag_cnt3 = rdata[2].tag.tag_cnt
        return oneData
      }
    }
  }
</script>


<style>
  .main{
    margin: 0 auto;
    width: 1600px;
    padding: 10px 10px;
    border: 1px solid #ccc;
  }
  .plist{
    height: auto;
    padding: 10px 10px;
    border: 1px solid #ccc;
  }
  .text {
    width: 200px;
    height: 400px;
    border-color: #ccc;
    margin-top: 10px;
  }
  .search{
    width: 200px;
  }
  .title{
    font-size: 20px;
    padding-bottom: 20px;
    border-bottom: 1px solid #ccc;
  }
 .el-form-item .el-form-item__content{
    margin-left: 2px;
  }
  .el-row {
    margin-bottom: 20px;
  &:last-child {
     margin-bottom: 0;
   }
  }
  .el-col {
    border-radius: 4px;
  }
  .bg-purple-dark {
    background: #99a9bf;
  }
  .bg-purple {
    background: #d3dce6;
  }
  .bg-purple-light {
    margin:  0 0 0 6px;
    background: #e5e9f2;
  }
  .grid-content {
    border-radius: 4px;
    min-height: 36px;
  }
  .row-bg {
    padding: 10px 0;
    background-color: #f9fafc;
  }
  .el-table .cell, .el-table th>div {
    padding-left: 4px;
    padding-right: 4px;
    box-sizing: border-box;
    text-overflow: ellipsis;
    text-align: left;
  }
  .footer a{
    color: #000;
    text-decoration: none;
  }
</style>
