import styled from "styled-components";

import UserInfo from "./Component/userInfo";
import TechStack from "./Component/techStack";
import axios from "axios";

const Profile = () => {
  const checkCookie = () => {
    // cookie가 없으면 5초 뒤 로그인 유도하는 moder 창
  };
  const getUserProfile = async () => {
    await axios.get("/user/profile");
  };
  return (
    <Container>
      <UserInfo></UserInfo>
      <TechStack></TechStack>
    </Container>
  );
};

export default Profile;

const Container = styled.div`
  padding: 30px;
  > div {
    padding: 30px 0px;
  }
`;
