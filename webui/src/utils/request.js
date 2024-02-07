import axios from "axios";
import { message } from "ant-design-vue";

const instance = axios.create({
  baseURL: "/apis",
  timeout: 10000,
  withCredentials: true,
});

export const get = async (url, params) => {
  let { data } = await instance.get(url, { params }).catch(catchError);
  return data;
};

function catchError(err) {
  if (err.response) {
    message.error(err.response.data.msg);
  } else if (err.request) {
    message.error("请求超时，请重试");
  } else {
    message.error(err.message);
  }
}

export const post = async (url, params) => {
  let { data } = await instance.post(url, params).catch(catchError);
  return data;
};
