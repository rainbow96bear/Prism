// AxiosConfig.ts

import axios, { AxiosInstance } from "axios";

const instance: AxiosInstance = axios.create({
  baseURL: "http://localhost/api", // baseURL을 원하는 주소로 설정
  // 여기에 다른 설정들도 추가할 수 있습니다.
});

export default instance;
