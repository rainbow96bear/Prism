import styled from "styled-components";
import { useEffect, useRef } from "react";
import { TitlePath } from "../GlobalType/TitlePath";
import { useNavigate } from "react-router-dom";

interface ComponentProps {
  setDropdown: React.Dispatch<React.SetStateAction<boolean>>;
  list: TitlePath[];
}

const DropDown: React.FC<ComponentProps> = ({ list, setDropdown }) => {
  const navigate = useNavigate();
  const move = (path: string) => {
    navigate(path);
    setDropdown(false); // 클릭한 후에 드롭다운을 닫습니다.
  };

  const dropdownRef = useRef<HTMLDivElement | null>(null);

  const handleClickOutside = (event: MouseEvent) => {
    if (
      dropdownRef.current &&
      !dropdownRef.current.contains(event.target as Node)
    ) {
      setDropdown(false);
    }
  };

  useEffect(() => {
    document.addEventListener("mousedown", handleClickOutside);

    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, []);

  return (
    <Box ref={dropdownRef}>
      {list.map((item, index) => (
        <Item
          key={"dropdown" + index}
          onClick={() => {
            if (item?.path != null) {
              move(item.path);
            } else if (item.func != null) {
              item.func();
            }
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
