import styled from "styled-components";

import UserInfo from "./Component/userInfo";
import TechStack from "./Component/techStack";

const Profile = () => {
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
