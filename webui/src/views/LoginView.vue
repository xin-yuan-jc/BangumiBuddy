<template>
  <a-flex align="center" class="background" justify="center">
      <div class="login">
        <div class="header">Login</div>
        <a-form :model="formState" class="form" name="basic" @finish="onFinish">
          <a-form-item
            :rules="[{ required: true, message: '请输入用户名' }]"
            name="username"
          >
            <a-input
              v-model:value="formState.username"
              :bordered="false"
              class="input"
              placeholder="Username"
            >
              <template #prefix>
                <UserOutlined style="color: rgba(0, 0, 0, 0.25)" />
              </template>
            </a-input>
          </a-form-item>

          <a-form-item
            :rules="[{ required: true, message: '请输入密码' }]"
            name="password"
          >
            <a-input-password
              v-model:value="formState.password"
              :bordered="false"
              class="input"
              placeholder="Password"
            >
              <template #prefix>
                <LockOutlined style="color: rgba(0, 0, 0, 0.25)" />
              </template>
            </a-input-password>
          </a-form-item>
          <a-form-item style="padding-top: 20px">
            <a-button class="button" html-type="submit" type="primary">登录</a-button>
          </a-form-item>
        </a-form>
      </div>
  </a-flex>
</template>
<script lang="ts" setup>
import { reactive } from 'vue'
import { login } from '@/apis/auth'
import { useRoute, useRouter } from 'vue-router'

interface LoginData {
  username: string;
  password: string;
}

const formState:LoginData = reactive({
  username: "",
  password: "",
});
const route = useRoute()
const router = useRouter()
const onFinish = async (values:LoginData) => {
  let { username, password } = values;
  const redirect = route.query.redirect as string || '/';
  await login(username, password)
  router.push(redirect)
};
</script>

<style lang="scss" scoped>
.background {
  height: 100vh;
  background-image: linear-gradient(to right, #c2f2fb, #a6c1ee);
  background-repeat: no-repeat;
}
.login {
  width: 358px;
  height: 433px;
  border-radius: 15px;
  background: linear-gradient(
      to right bottom,
      rgba(255, 255, 255, 0.7),
      rgba(255, 255, 255, 0.5),
      rgba(255, 255, 255, 0.4)
  );
  backdrop-filter: blur(10px);
  box-shadow: 0 0 20px #a29bfe;
}
.header {
  font-size: 38px;
  font-weight: bold;
  text-align: center;
  padding-top: 80px;
  padding-bottom: 95px;
}
.form {
  width: 100%;
  padding-right: 20px;
  padding-left: 20px;
}
.input {
  border-bottom: 1px solid rgb(128, 125, 125);
}
.button {
  width: 100%;
  background-image: linear-gradient(to right, #a6c1ee, #fbc2eb);
}
</style>
