import axios from "axios";
import { useNavigate } from "react-router-dom";
import styled from "styled-components";

const FuncBar = () => {
  const navigate = useNavigate();
  return (
    <Box>
      <div
        onClick={() => {
          navigate("/admin/setting/tech");
        }}>
        기술 스택 관리
      </div>
    </Box>
  );
};

export default FuncBar;

const Box = styled.div`
  display: flex;
  align-items: center;
  > div {
    margin: 20px;
    font-size: 1.2rem;
  }
`;
