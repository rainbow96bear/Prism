import { useState } from "react";
import styled from "styled-components";

interface HashTagInputProps {
  prevHashTag: string[];
  setHashTag: React.Dispatch<React.SetStateAction<string[]>>;
}

const HashTagInput: React.FC<HashTagInputProps> = ({
  prevHashTag,
  setHashTag,
}) => {
  const [inputValue, setInputValue] = useState("");

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(event.target.value);
  };

  const handleInputKeyDown = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === "Enter" && inputValue.trim() !== "") {
      // 이미 존재하는 해시태그인지 검사
      if (!prevHashTag.includes(inputValue.trim())) {
        setHashTag((prevHashTag) => [...prevHashTag, inputValue.trim()]);
        setInputValue("");
      }
    }
  };

  const handleInputBlur = () => {
    if (inputValue.trim() !== "") {
      // 이미 존재하는 해시태그인지 검사
      if (!prevHashTag.includes(inputValue.trim())) {
        setHashTag((prevHashTag) => [...prevHashTag, inputValue.trim()]);
        setInputValue("");
      }
    }
  };

  return (
    <Container>
      <HashTagLogo>#</HashTagLogo>
      <InputBox
        type="text"
        placeholder="태그 입력"
        maxLength={10}
        value={inputValue}
        onChange={handleInputChange}
        onKeyDown={handleInputKeyDown}
        onBlur={handleInputBlur}
      />
    </Container>
  );
};

export default HashTagInput;

const Container = styled.div`
  display: flex;
  align-items: center;
  font-size: 1.2rem;
`;

const HashTagLogo = styled.div`
  padding: 5px;
  font-weight: bold;
`;

const InputBox = styled.input`
  border: none;
  padding: 5px;
  outline: none;
  font-weight: bold;
`;
