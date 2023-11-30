import styled from "styled-components";

import Logo from "./Components/logo";
import Category from "./Components/category";

const Header = () => {
  return (
    <Box>
      <Left>
        <S_Logo></S_Logo>
        <S_Category></S_Category>
      </Left>
      <Right></Right>
    </Box>
  );
};

export default Header;

const Box = styled.div`
  display: flex;
  height: 55px;
  padding: 0px 10px;
  border-bottom: 2px solid gray;
  div {
    padding: 0px 10px;
    display: flex;
  }
`;

const Left = styled.div``;
const Right = styled.div``;

const S_Logo = styled(Logo)``;
const S_Category = styled(Category)``;
