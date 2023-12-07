import styled from "styled-components";
import { useLocation, useNavigate } from "react-router-dom";

interface CommunityProps {
  iscommunity: string;
}

interface ProjectProps {
  isproject: string;
}

const Category: React.FC = () => {
  const location = useLocation();
  const navigate = useNavigate();

  const iscommunity: string = String(location.pathname === "/community");
  const isproject: string = String(location.pathname === "/project");

  const handleCommunityClick = () => {
    navigate("/community");
  };
  const handleProjectClick = () => {
    navigate("/project");
  };
  return (
    <Box>
      <Community onClick={handleCommunityClick} iscommunity={iscommunity}>
        커뮤니티
      </Community>
      <Project onClick={handleProjectClick} isproject={isproject}>
        프로젝트
      </Project>
    </Box>
  );
};

const Box = styled.div`
  display: flex;
  height: 100%;
  div {
    margin: 0px 10px 0px 10px;
    display: flex;
    align-items: center;
    cursor: pointer;
    font-size: 1.3rem;
  }
`;

const Community = styled.div<CommunityProps>`
  color: ${({ iscommunity }) =>
    iscommunity === "true" ? "black" : "lightgray"};
  border-bottom: ${({ iscommunity }) =>
    iscommunity === "true" ? "4px solid green" : ""};
`;

const Project = styled.div<ProjectProps>`
  color: ${({ isproject }) => (isproject === "true" ? "black" : "lightgray")};
  border-bottom: ${({ isproject }) =>
    isproject === "true" ? "4px solid green" : ""};
`;

export default Category;
