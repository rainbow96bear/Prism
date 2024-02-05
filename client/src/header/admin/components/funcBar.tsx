import { useSelector } from "react-redux";
import { useNavigate } from "react-router-dom";
import styled from "styled-components";
import { RootState } from "../../../app/store";

const FuncBar = () => {
  const admin_info = useSelector(
    (state: RootState) => state.adminReducer.admin_info
  );
  const navigate = useNavigate();
  return (
    <Box>
      {admin_info.id != "" ? (
        <div
          onClick={() => {
            navigate("/admin/setting/tech");
          }}>
          기술 스택 관리
        </div>
      ) : (
        <></>
      )}
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
