<template>
    <div class="header">
        <img src="../assets/IMT_Atlantique_logo_RVB.svg" alt="logo"/>
        <span class="header_text" @click="changePage">
            Simple Forum
        </span>
        <span style="float: right" class="header_text">
            <el-button type="text" @click="dialogVisible = true" style="padding-top: 0">Post</el-button>
        </span>
        <el-dialog
                title="Post"
                :visible.sync="dialogVisible"
                width="30%"
                :before-close="handleClose">
            <el-form ref="form" :model="form" label-width="80px">
                <el-form-item label="Title">
                    <el-input v-model="form.title"></el-input>
                </el-form-item>
                <el-form-item label="Content">
                    <el-input type="textarea" v-model="form.desc"></el-input>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button @click="closePost">Cancer</el-button>
                <el-button type="primary" @click="closePost">Confirm</el-button>
            </span>
        </el-dialog>
    </div>
</template>

<script>
    export default {
        name: "Header",
        data() {
            return {
                dialogVisible: false,
                form: {
                    title: '',
                    desc: ''
                }
            };
        },
        methods: {
            emptyPost() {
                this.form.title = '';
                this.form.desc = '';
            },
            closePost(submit) {
                this.dialogVisible = false;
                if (submit) {
                    let titleSend = this.form.title;
                    let descSend = this.form.desc;
                    if (titleSend && descSend){
                        let postObj = {
                            title: titleSend,
                            desc: descSend,
                            author: (new Date()).valueOf()
                        };
                        let config = {
                            header: {
                                'Content-Type':'application/json'
                            }
                        };
                        const axiosAjax = this.axios.create({
                            timeout: 1000*60,
                            withCredentials: true
                        });
                        axiosAjax.post('http://localhost:8080/post', postObj, config).then((res)=>{
                            console.log(res.data)
                            this.$notify({
                                title: "submit success",
                                type: "success",
                                message: "thanks !",
                                duration: 2000
                            });
                            this.emptyPost();
                        }).catch((err)=>{
                            console.log(err);
                            this.emptyPost();
                        });
                    } else {
                        this.$notify({
                            title: "submit error",
                            type: "error",
                            message: "error",
                            duration: 2000
                        })
                    }
                }
            },
            handleClose(done) {
                done()
                this.emptyPost();
            },
            changePage() {
                this.$router.push("/primary")
            }
        }
    }
</script>

<style scoped>
    .header {
        height: 100%;
        width: 100%;
        align-items: center;
        padding-right: 2vw;
        box-sizing: border-box;
    }

    img {
        height: 100%;
        width: 10vw;
        float: left;
    }

    span.header_text {
        text-align: left;
        display: inline-block;
        float: left;
        align-items: center;
        padding-top: 2vw;
    }

</style>