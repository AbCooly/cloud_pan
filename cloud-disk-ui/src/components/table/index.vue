<template>
  <div style="height: 100%">
    <el-table
        :data="fileList"
        style="width: 100%">
      <el-table-column
          prop="name"
          width="260"
      >
        <template slot-scope="scope">
          <span v-if="scope.row.ext">{{ scope.row.name }}</span>
          <el-button v-else type="text" @click="endDir(scope.row)">{{ scope.row.name }}</el-button>
        </template>
      </el-table-column>
      <el-table-column
          label="类型"
          width="180">
        <template slot-scope="scope">
          <span>{{ scope.row.ext|getFileType }}</span>
        </template>
      </el-table-column>
      <el-table-column
          label="操作"
          width="200">
        <template slot-scope="scope">
          <el-button v-if="scope.row.ext" @click="downloadFile(scope.row)" type="text" size="small">下载</el-button>
          <el-button @click="edit(scope.row)" type="text" size="small">修改</el-button>
          <!--          <el-button @click="shareFile(scope.row)" type="text" size="small">分享</el-button>-->
          <el-button @click="deleteFile(scope.row)" type="text" size="small">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <div style="width: 100%;text-align: center">
      <el-pagination
          @current-change="pageChange"
          :page-size="size"
          layout="prev, pager, next"
          :total="count">
      </el-pagination>
    </div>
  </div>


</template>

<script>
import {getFileType} from "@/utils/file_type";
import {deleteFile} from "@/api/user";

export default {
  //过滤器
  filters: {
    getFileType(item) {
      return getFileType(item)
    }
  },
  props: {
    fileList: {
      type: Array,
      default() {
        return [{
          ext: "",
          id: 0,
          name: "",
          path: "",
          repositoryIdentity: "",
          message: "",
        }]
      }
    },
    count: {
      type: Number,
      default() {
        return 0;
      }
    },
    size: {
      type: Number,
      default() {
        return 10
      }
    }
  },
  data() {
    return {}
  },
  methods: {
    pageChange(e) {
      this.$emit('change-page', e)
    },
    //进入文件夹
    endDir(e) {
      this.$emit('change-dir', e)
    },
    changeBlob(url) {
      return new Promise(resolve => {
        const xhr = new XMLHttpRequest();
        xhr.open('GET', url, true);
        xhr.responseType = 'blob';
        xhr.onload = () => {
          if (xhr.status === 200) {
            resolve(xhr.response);
          }
        };
        xhr.send();
      });
    },
    //下载文件
    saveAs(blob, filename) {
      let link = document.createElement('a');
      link.href = window.URL.createObjectURL(blob);
      link.download = filename;
      link.click();
    },
    //下载
    downloadFile(e) {
      this.changeBlob(e.path).then(res => {
        this.saveAs(res, e.name)
      })

    },
    //移动
    moved(e) {

    },
    //分享
    shareFile(e) {
    },
    //删除
    deleteFile(e) {
      this.$confirm('此操作将永久删除该文件, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        deleteFile({id: e.id}).then(res => {
          if (res.data.code === 200) {
            if (e.ext === "") {
              this.$emit("change-go-back")
            } else {
              this.$emit("change-list")
            }
          }
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        });
      });
    }
  }
}
</script>
