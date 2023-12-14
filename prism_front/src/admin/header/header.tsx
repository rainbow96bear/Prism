import styled from "styled-components";
import Logo from "./components/logo";

const AdminHeader = () => {
  return (
    <Box>
      <Group className="CategoryGroup">
        <Logo></Logo>
      </Group>
    </Box>
  );
};

export default AdminHeader;

const Box = styled.div`
  display: flex;
  justify-content: space-between;

  width: 90%;
  height: 100%;
  .CategoryGroup {
    display: flex;
    font-weight: bold;
    flex: 1;
    min-width: 277px;
  }
  .SearchGroup {
    display: flex;
    justify-content: center;
    align-items: center;
    flex: 2;
    min-width: 277px;
  }
  .FuncBarGroup {
    flex: 1;
    min-width: 277px;
  }
`;

const Group = styled.div``;
