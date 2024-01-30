import axios from "./../../../configs/AxiosConfig";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import styled from "styled-components";

interface setAdmin_info {
  setAdmin_info: Function;
}

const Root: React.FC<setAdmin_info> = ({ setAdmin_info }) => {
  const navigate = useNavigate();
  const [password, setPassword] = useState(""); // 입력한 비밀번호를 상태로 관리

  const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(e.target.value);
  };

  const handleSubmit = async () => {
    try {
      const loginResult = (
        await axios.post(
          "/admin/login",
          { password: password },
          {
            withCredentials: true,
          }
        )
      ).data;
      if (loginResult?.admin_info.id != "") {
        setAdmin_info(loginResult?.admin_info);
        navigate(`/admin/home/${loginResult?.admin_info.id}`);
      } else {
        alert("접근 번호를 정확히 입력하세요.");
      }
    } catch (error) {
      console.error("Admin 정보를 가져오는 중 에러 발생:", error);
    }
  };

  return (
    <Box>
      <DescribeBox>관리자 접속</DescribeBox>
      <PasswordBox>
        <input
          type="password"
          value={password}
          onChange={handlePasswordChange}
        />
        <button onClick={handleSubmit}>접속</button>
      </PasswordBox>
    </Box>
  );
};

export default Root;

const Box = styled.div``;

const DescribeBox = styled.div`
  font-size: 20px;
  font-weight: bold;
`;

const PasswordBox = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  input {
    padding: 0px;
    margin: 2px;
    font-size: 25px;
  }
  button {
    padding: 0px;
    margin: 2px;
    font-size: 20px;
    font-weight: bold;
  }
`;
