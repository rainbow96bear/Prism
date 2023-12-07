import styled from "styled-components";
import { FC } from "react";

interface TechItemProps {
  tech: string;
  level: number;
  maxLevel: number;
}

const TechItem: FC<TechItemProps> = ({ tech, level, maxLevel }) => {
  return (
    <Container>
      <Title>{tech}</Title>
      <LevelBar>
        {Array.from({ length: maxLevel }, (_, index) => (
          <LevelBlock key={index} filled={index < level} />
        ))}
      </LevelBar>
    </Container>
  );
};

export default TechItem;

interface LevelBlockProps {
  filled: boolean;
}

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
`;

const LevelBar = styled.div`
  display: flex;
`;

const LevelBlock = styled.div<LevelBlockProps>`
  width: 20px;
  height: 20px;
  background-color: ${(props) =>
    props.filled ? "hsl(999, 70%, 50%)" : "transparent"};
  border: 2px solid hsl(999, 70%, 50%);
  margin-right: 5px;
`;
