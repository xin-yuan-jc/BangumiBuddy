<template>
  <a-flex class="background" justify="center" align="center">
    <a-flex class="login">
      <a-space direction="vertical" style="width: 100%">
        <br />
        <br />
        <br />
        <div class="header">Login</div>
        <br />
        <br />
        <a-form :model="formState" name="basic" @finish="onFinish" class="form">
          <a-form-item
            name="username"
            :rules="[{ required: true, message: '请输入用户名' }]"
          >
            <a-input
              v-model:value="formState.username"
              placeholder="Username"
              :bordered="false"
              class="input"
            >
              <template #prefix>
                <UserOutlined style="color: rgba(0, 0, 0, 0.25)" />
              </template>
            </a-input>
          </a-form-item>

          <a-form-item
            name="password"
            :rules="[{ required: true, message: '请输入密码' }]"
          >
            <a-input-password
              v-model:value="formState.password"
              placeholder="Password"
              :bordered="false"
              class="input"
            >
              <template #prefix>
                <LockOutlined style="color: rgba(0, 0, 0, 0.25)" />
              </template>
            </a-input-password>
          </a-form-item>
          <br />
          <a-form-item>
            <a-button type="primary" html-type="submit" class="button"
              >登录</a-button
            >
          </a-form-item>
        </a-form>
      </a-space>
    </a-flex>
  </a-flex>
</template>
<script setup>
import { reactive } from "vue";
import { UserOutlined, LockOutlined } from "@ant-design/icons-vue";
import { login } from "@/api/auth";

const formState = reactive({
  username: "",
  password: "",
});
const onFinish = (values) => {
  let { username, password } = values;
  login({ username, password });
};
</script>

<style lang="scss" scoped>
.background {
  height: 100vh;
  background-image: linear-gradient(to right, #fbc2eb, #a6c1ee);
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
}
.form {
  width: 100%;
  padding: 20px;
}
.input {
  border-bottom: 1px solid rgb(128, 125, 125);
}
.button {
  width: 100%;
  background-image: linear-gradient(to right, #a6c1ee, #fbc2eb);
}
</style>
