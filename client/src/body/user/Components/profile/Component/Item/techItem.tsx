import styled from "styled-components";
import { FC } from "react";
import { TechData } from "../../../../../../GlobalType/Tech";

const TechItem: FC<TechData> = ({ tech_name, level }) => {
  const maxLevel = 10;
  return (
    <Container>
      <Title>{tech_name}</Title>
      <LevelBar>
        {Array.from({ length: maxLevel }, (_, index) =>
          index < level ? (
            <LevelFillBlock key={index} />
          ) : (
            <LevelEmptyBlock key={index} />
          )
        )}
      </LevelBar>
    </Container>
  );
};

export default TechItem;

const Container = styled.div`
  display: flex;
  padding: 30px 0px;
  > div {
    padding: 10px 15px;
  }
`;

const Title = styled.div`
  background-color: hsl(0, 0%, 90%);
  border: 2px solid lightgray;
  border-radius: 20px;
  font-weight: bold;
  width: 110px;
  text-align: center;
`;

const LevelBar = styled.div`
  display: flex;
`;

const LevelEmptyBlock = styled.div`
  width: 20px;
  height: 20px;
  background-color: transparent;
  border: 2px solid hsl(999, 70%, 50%);
  margin-right: 5px;
`;

const LevelFillBlock = styled.div`
  width: 20px;
  height: 20px;
  background-color: hsl(999, 70%, 50%);
  border: 2px solid hsl(999, 70%, 50%);
  margin-right: 5px;
`;
