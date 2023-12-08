import styled from "styled-components";
import { useEffect, useState } from "react";

import BeforeLogin from "./beforeLogin";
import AfterLogin from "./afterLogin";
import axios from "axios";

interface UserInfo {
  userID: string;
  imgUrl: string;
}

const FuncBar: React.FC = () => {
  const [userInfo, setUserInfo] = useState<UserInfo | null>(null);

  const getUserInfo = async () => {
    try {
      const result = await axios.get("http://localhost:8080/userinfo", {
        withCredentials: true,
      });

      const { sub, picture } = result.data;
      setUserInfo({ userID: sub, imgUrl: picture });
    } catch (error) {
      console.error("Error fetching user info:", error);
    }
  };

  useEffect(() => {
    const User_Login = document.cookie;
    const hasCookie = User_Login.includes("user_login");
    console.log(hasCookie);
    if (hasCookie) {
      getUserInfo();
    } else {
      setUserInfo(null);
    }
    console.log(userInfo);
  }, []);

  return (
    <Box>
      {userInfo != null ? (
        <AfterLogin userID={userInfo?.userID} imgUrl={userInfo?.imgUrl} />
      ) : (
        <BeforeLogin />
      )}
    </Box>
  );
};

export default FuncBar;

const Box = styled.div`
  display: flex;
  justify-content: right;
  height: 100%;
  align-items: center;
`;
