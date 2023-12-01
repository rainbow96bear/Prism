import styled from "styled-components";

import Logo from "./Components/logo";
import Category from "./Components/category";
import SearchBox from "./Components/searchBox";
import FuncBar from "./Components/funcBar";

const Header = () => {
  return (
    <Box>
      <Group className="CategoryGroup">
        <Logo></Logo>
        <Category></Category>
      </Group>
      <Group className="SearchGroup">
        <SearchBox></SearchBox>
      </Group>
      <Group className="FuncBarGroup">
        <FuncBar></FuncBar>
      </Group>
    </Box>
  );
};

export default Header;

const Box = styled.div`
  display: flex;
  justify-content: space-between;
  height: 100%;
  .CategoryGroup {
    display: flex;
    font-weight: bold;
    width: 25%;
    min-width: 265px;
  }
  .SearchGroup {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 50%;
    min-width: 265px;
  }
  .FuncBarGroup {
    width: 25%;
    min-width: 265px;
  }
`;

const Group = styled.div``;
