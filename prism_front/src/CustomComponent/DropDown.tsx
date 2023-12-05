import styled from "styled-components";

import { TitlePath } from "../GlobalType/TitlePath";
import { useNavigate } from "react-router-dom";

interface ComponentProps {
  list: TitlePath[];
}

const DropDown: React.FC<ComponentProps> = ({ list }) => {
  const navigate = useNavigate();
  const move = (path: string) => {
    navigate(path);
  };
  return (
    <Box>
      {list.map((item, index) => (
        <Item
          key={"dropdown" + index}
          onClick={() => {
            move(item.path);
          }}>
          {item.title}
        </Item>
      ))}
    </Box>
  );
};

export default DropDown;

const Box = styled.div`
  // border: 2px solid lightgray;
  height: fit-content;
`;

const Item = styled.ul`
  liststyle: "none";
  width: 100%;
  padding: 10px 0px;
  margin: 0px;
  display: flex;
  justify-content: center;
  border: 1px solid lightgray;
  font-weight: bold;
  font-size: 1.2rem;
`;
