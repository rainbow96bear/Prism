import styled from "styled-components";
import { useDispatch } from "react-redux";
import TechItem from "./Item/techItem";
import { AppDispatch, RootState } from "../../../../app/store";
import { useSelector } from "react-redux";
import { useNavigate, useParams } from "react-router";
import { useEffect } from "react";
import { getTechList } from "../../../../app/slices/profile/tech_data";
const TechStack = () => {
  const { id } = useParams();

  const user = useSelector((state: RootState) => state.user);

  const dispatch = useDispatch<AppDispatch>();
  const techList = useSelector((state: RootState) => state.tech_data);

  const navigator = useNavigate();

  useEffect(() => {
    dispatch(getTechList(id));
  }, []);
  return (
    <Container>
      <TitleBox>
        <Title>기술 스택</Title>
        {user.user_id == id && (
          <button
            onClick={() => {
              navigator("/profile/update/techlist");
            }}>
            수정
          </button>
        )}
      </TitleBox>
      {techList.tech_list?.map((value, index) => (
        <TechItem
          key={index}
          tech_name={value.tech_name}
          level={value.level}></TechItem>
      ))}
    </Container>
  );
};

export default TechStack;

const Container = styled.div``;

const TitleBox = styled.div`
  display: flex;
  justify-content: space-between;
`;

const Title = styled.div`
  font-weight: bold;
  font-size: 1.5rem;
`;
