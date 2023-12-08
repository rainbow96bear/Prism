import styled from "styled-components";
import TechItem from "./Item/techItem";

const TechStack = () => {
  const techProps = {
    tech: "Golang",
    level: 6,
    maxLevel: 10,
  };

  return (
    <Container>
      <Title>기술 스택</Title>
      <TechItem {...techProps} />
    </Container>
  );
};

export default TechStack;

const Container = styled.div``;

const Title = styled.div`
  font-weight: bold;
  font-size: 1.5rem;
`;
