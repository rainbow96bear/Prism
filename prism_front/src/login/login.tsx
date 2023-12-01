import styled from "styled-components";
import { useEffect } from "react";
import axios from "axios";

const Login = () => {
  useEffect(() => {
    const fetchData = async () => {
      try {
        const code = new URL(window.location.href).searchParams.get("code");
        const result = await axios.get(
          `http://localhost:8080/kakaoLogin?code=${code}`
        );
        console.log(result);
      } catch (error) {
        console.error("Error fetching data:", error);
        // 에러 처리 로직 추가 가능
      }
    };

    fetchData();
  }, []);
  return <div>Loading...</div>;
};

export default Login;
