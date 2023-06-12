<template>
  <div class="main-content">
    <div class="nav-class">
      <div class="nav-class-user">
        <span style="font-weight: bold">欢迎用户：</span>
        <span>{{ userInfo.name }}</span>
        <span style="font-weight: bold">当前时间：</span>
        <span style="color: blue">{{ nowDate }}</span>
      </div>
      <div class="nav-class-system">
        <el-button type="text" slot="label"><i class="el-icon-user"></i> 个人资料</el-button>
        <el-button type="text" slot="label"><i class="el-icon-key"></i> 修改密码</el-button>
        <el-button type="text" slot="label" @click="exitClick"><i class="el-icon-switch-button"></i> 退出登录</el-button>
      </div>
    </div>
    <el-divider></el-divider>
    <el-page-header style="margin-left: 20px" @back="goBack" :content="perObj.name">
    </el-page-header>
    <el-divider></el-divider>
    <div class="new-class-button">
      <el-button type="primary" size="small" slot="label" @click="makeDir"><i class="el-icon-folder"></i>新建文件夹
      </el-button>
      <el-button type="success" size="small" slot="label" @click="uploadFile"><i class="el-icon-upload"></i>上传文件
      </el-button>
    </div>
    <div class="table-class">
      <el-tabs tab-position="left" v-model="queryParam.type" style="height: 600px;" @tab-click="handleTabClick">
        <el-tab-pane name="all" label="all">
          <span slot="label"><i class="el-icon-folder-opened"></i> 全部</span>
          <table-data @change-page="getPageFun" @change-dir="changeDir" @change-list="getUserFileList"
                      @change-go-back="goBack" :fileList="fileList" :count="pageCount"
                      :size="queryParam.size"></table-data>
        </el-tab-pane>
        <el-tab-pane name="image" label="images">
          <span slot="label"><i class="el-icon-picture"></i> 图片</span>
          <table-data @change-page="getPageFun" :fileList="fileList" @change-list="getUserFileList" :count="pageCount"
                      :size="queryParam.size"></table-data>
        </el-tab-pane>
        <el-tab-pane name="doc" label="doc">
          <span slot="label"><i class="el-icon-document"></i> 文档
          </span>
          <table-data @change-page="getPageFun" :fileList="fileList" @change-list="getUserFileList" :count="pageCount"
                      :size="queryParam.size"></table-data>
        </el-tab-pane>
        <el-tab-pane name="videos" label="videos">
          <span slot="label"><i class="el-icon-video-play"></i> 视频</span>
          <table-data @change-page="getPageFun" :fileList="fileList" @change-list="getUserFileList" :count="pageCount"
                      :size="queryParam.size"></table-data>
        </el-tab-pane>
        <el-tab-pane name="other" label="other">
          <span slot="label"><i class="el-icon-more"></i> 其他</span>
          <table-data @change-page="getPageFun" :fileList="fileList" @change-list="getUserFileList" :count="pageCount"
                      :size="queryParam.size"></table-data>
        </el-tab-pane>
      </el-tabs>
    </div>
    <el-dialog
        title="提示"
        :visible.sync="uploadDialog"
        width="450px"
        :before-close="closeDialog">
      <el-upload
          drag
          :show-file-list="false"
          :limit="10"
          :auto-upload="true"
          action=""
          :http-request="beforeUpload">
        <i class="el-icon-upload"></i>
        <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
        <div class="el-upload__tip" slot="tip">一次最多只能上传10个文件</div>
      </el-upload>

      <span slot="footer" class="dialog-footer">
    <el-button @click="uploadDialog = false">取 消</el-button>
    <el-button type="primary" @click="submitUpload">确 定</el-button>
  </span>
    </el-dialog>
  </div>
</template>

<script>
import dayjs from 'dayjs'
import axios from 'axios'
import TableData from "@/components/table/index"
import {getUserById, getUserFileList, makeCreateDir, userDetail,userFileLink} from "@/api/user"
import router from "@/router";

var token = localStorage.getItem('token')
require('dayjs/locale/zh-cn')
export default {
  components: {
    TableData
  },
  data() {
    return {
      uploadFileList: [],
      uploadDialog: false,
      nowDate: "",
      nowTimer: "",
      userInfo: {
        userName: ""
      },
      handerContent: "",
      pageCount: 0,
      queryParam: {
        type: "all",
        id: 0,
        page: 1,
        size: 10
      },
      //当前进入的文件夹
      perObj: {},
      fileList: [
        {
          ext: "",
          id: 0,
          name: "",
          path: "",
          repositoryIdentity: "",
          message: "",
        }
      ]
    }
  },
  methods: {
    //进行最终保存到用户池
    submitUpload() {
      let param = {
        parentId: 0,
        repositoryIdentity: "",
        ext: "",
        name: "",
        type: ""
      }
      for (const file of this.uploadFileList) {
        param.parentId = this.queryParam.id
        param.repositoryIdentity = file.id
        param.ext = file.ext
        param.name = file.name
        param.type = file.type
        userFileLink(param).then(res=>{
          this.getUserFileList()
        })
      }
      this.uploadFileList = []
      this.uploadDialog = false
    },
    //文件上传前准备
    beforeUpload(e) {
      let formData = new FormData();
      formData.append("file", e.file)
      formData.append("parentId",  this.queryParam.id)

      axios({
        method: 'post',
        url: '/api/file/upload',
        data: formData,
        headers: {
          'Content-Type': 'multipart/form-data',
          'Authorization': localStorage.getItem("token")
        }
      }).then(res => {
        if (res.data.code === 200) {
          this.uploadFileList.push(res.data.data)
        }
      })
    },
    //获取用户数据
    getUserInfo() {
      userDetail().then(res => {
        if (res.data.code === 200) {
          localStorage.setItem("user", res.data.data)
          this.userInfo = res.data.data
        }
      })
    },
    //获取用户文件列表
    getUserFileList() {
      getUserFileList(this.queryParam).then(res => {
        this.fileList = res.data.data.list
        this.pageCount = res.data.data.count
      })
    },

    //获取子组件页面变更
    getPageFun(e) {
      this.queryParam.page = e
      this.getUserFileList()
    },
    //获取子组件传递的文件夹变更数据
    changeDir(e) {
      this.perObj = e
      this.queryParam.id = e.id
      this.queryParam.size = 10
      this.queryParam.page = 1
      this.getUserFileList()
    },
    handleSelect(key, keyPath) {
      console.log(key, keyPath);
    },
    //返回上一页
    goBack() {
      if (!this.perObj.parentId) {
        this.queryParam.id = 0
      } else {
        this.queryParam.id = this.perObj.parentId
      }
      this.queryParam.size = 10
      this.queryParam.page = 1
      this.getUserFileList()
      if (this.perObj.parentId === 0 || !this.perObj.parentId) {
        this.perObj = {}
      } else {
        getUserById({id: this.perObj.parentId}).then(res => {
          this.perObj = res.data.data
        })
      }
    },

    //切换tag
    handleTabClick(tab, event) {
      this.queryParam.type = tab.name
      this.queryParam.page = 1
      this.getUserFileList()
    },
    //退出登录
    exitClick() {
      this.$confirm('其否确认退出登录？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        localStorage.removeItem("token")
        localStorage.removeItem("user")
        router.replace({
          path: "/"
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消退出'
        });
      });
    },
    //新建文件夹
    makeDir() {
      let param = {
        parentId: this.queryParam.id,
        name: ""
      }
      this.$prompt('请输入文件夹名称', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        inputValidator(value) {
          return value !== "" && value !== null && value.length > 0
        },
        inputErrorMessage: '请输入文件夹名称'
      }).then(({value}) => {
        param.name = value
        makeCreateDir(param).then(res => {
          this.getUserFileList()
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '取消输入'
        });
      });

    },
    //文件上传
    uploadFile() {
      this.uploadDialog = true
    },
    closeDialog() {
      this.uploadDialog = false
    }
  },
  created() {
    this.getUserFileList()
    this.getUserInfo()
    this.nowTimer = setInterval(() => {
      this.nowDate = dayjs().locale('zh-cn').format('YYYY-MM-DD HH:mm:ss dddd');
    }, 1000);
  },
  beforeDestroy() {
    //销毁定时
    if (this.nowTimer) {
      clearInterval(this.nowTimer);
    }
  }
}
</script>

<style lang="less" scoped>
.main-content {
  height: 100%;

  .nav-class {
    height: 50px;

    .nav-class-system {
      position: absolute;
      top: 25px;
      right: 10px;
    }

    .nav-class-user {
      position: absolute;
      top: 25px;
      left: 20px;
    }

    span {
      margin-left: 10px;
    }
  }

  .new-class-button {
    float: right;
  }

  .table-class {
    height: 90%;
    width: 100%;
  }
}
</style>
