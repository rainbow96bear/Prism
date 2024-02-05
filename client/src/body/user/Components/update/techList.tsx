import React, { useState, useEffect, useRef, KeyboardEvent } from "react";
import { useDispatch } from "react-redux";
import { AppDispatch, RootState } from "../../../../app/store";
import { useSelector } from "react-redux";
import { getUserTechList } from "../../../../app/slices/profile/tech_data";
import TechItem from "./TechListComponent/techItem";
import styled from "styled-components";
import { TechData, TechInfo } from "../../../../GlobalType/Tech";
import { useNavigate } from "react-router";
import axios from "../../../../configs/AxiosConfig";

const TechList = () => {
  const dispatch = useDispatch<AppDispatch>();
  const user_tech_list = useSelector((state: RootState) => state.tech_data);
  const user = useSelector((state: RootState) => state.user);
  const navigator = useNavigate();
  const [userTechList, setUserTechList] = useState<TechData[]>(
    user_tech_list.tech_list ?? []
  );

  const [techList, setTechList] = useState<TechInfo[]>([]);
  const [searchTerm, setSearchTerm] = useState<string>("");
  const [searchResults, setSearchResults] = useState<string[]>([]);
  const [selectedResultIndex, setSelectedResultIndex] = useState(-1);
  const [showResults, setShowResults] = useState(false);
  const containerRef = useRef<HTMLDivElement>(null);

  // 기존 기술 목록을 가져오기
  useEffect(() => {
    dispatch(getUserTechList(user.id));
  }, [dispatch, user.id]);

  useEffect(() => {
    const getTechList = async () => {
      try {
        const result = await axios.get("/users/profiles/techlist", {
          withCredentials: true,
        });
        // name과 count가 같이 전달되기에 name만 뽑아쓰도록 하기
        setTechList(result.data);
      } catch (error) {
        console.error("techList 받아오는 중 에러 발생:", error);
      }
    };
    getTechList();
  }, []);

  // 검색어 입력 시 호출되는 함수
  const handleSearchChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const { value } = event.target;
    setSearchTerm(value);

    // 검색어에 따라 결과 필터링
    const filteredResults = techList
      .filter((tech) =>
        tech?.tech_name?.toLowerCase().includes(value.toLowerCase())
      )
      .map((tech) => tech.tech_name || ""); // 추출된 기술명들로 이루어진 배열을 생성

    setSearchResults(filteredResults);
    setSelectedResultIndex(-1);
    setShowResults(!!value);
  };

  // 기술 추가 함수
  const handleAddTech = (selectedTech: string) => {
    setSearchTerm("");

    // Check for duplicate values
    if (userTechList?.some((tech) => tech.tech_name === selectedTech)) {
      alert("중복된 값은 추가할 수 없습니다.");
      return;
    }

    // 선택한 결과를 기존 기술 목록에 추가 (레벨 0으로)
    setUserTechList((prevTechList) => [
      ...prevTechList,
      { tech_name: selectedTech, level: 0 },
    ]);
    setShowResults(false);
  };
  // 외부를 클릭했을 때 결과 숨기기
  const handleClickOutside = (event: MouseEvent) => {
    if (
      containerRef.current &&
      !containerRef.current.contains(event.target as Node)
    ) {
      setShowResults(false);
    }
  };

  // 키보드 이벤트 처리
  const handleKeyDown = (event: KeyboardEvent<HTMLInputElement>) => {
    if (event.key === "ArrowUp") {
      event.preventDefault();
      setSelectedResultIndex((prevIndex) =>
        prevIndex > 0 ? prevIndex - 1 : searchResults.length - 1
      );
    } else if (event.key === "ArrowDown") {
      event.preventDefault();
      setSelectedResultIndex((prevIndex) =>
        prevIndex < searchResults.length - 1 ? prevIndex + 1 : 0
      );
    } else if (event.key === "Enter" && selectedResultIndex !== -1) {
      handleAddTech(searchResults[selectedResultIndex]);
    }
  };

  // 기존 기술 삭제 함수
  const handleRemoveTech = (techName: string) => {
    setUserTechList((prevTechList) =>
      prevTechList.filter((tech) => tech.tech_name !== techName)
    );
  };

  // 기술 레벨 낮추기 함수
  const handleDecrementLevel = (techName: string) => {
    setUserTechList((prevTechList) =>
      prevTechList.map((tech) =>
        tech.tech_name === techName && tech.level > 0
          ? { ...tech, level: tech.level - 1 }
          : tech
      )
    );
  };

  // 기술 레벨 올리기 함수
  const handleIncrementLevel = (techName: string) => {
    setUserTechList((prevTechList) =>
      prevTechList.map((tech) =>
        tech.tech_name === techName && tech.level < 10
          ? { ...tech, level: tech.level + 1 }
          : tech
      )
    );
  };

  const saveTech = async (userTechList: TechData[]) => {
    await axios.put(
      `/users/profiles/${user.id}/techs`,

      userTechList,

      {
        withCredentials: true,
      }
    );
    navigator(`/profile/${user.id}`);
  };

  // 외부 클릭 이벤트 리스너 등록
  useEffect(() => {
    document.addEventListener("mousedown", handleClickOutside);
    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, []);

  return (
    <Container ref={containerRef}>
      {userTechList?.map((value, index) => (
        <TechItem
          key={index}
          tech_name={value.tech_name}
          level={value.level}
          onRemove={() => handleRemoveTech(value.tech_name)}
          onDecrement={() => handleDecrementLevel(value.tech_name)}
          onIncrement={() => handleIncrementLevel(value.tech_name)}
        />
      ))}
      <SearchContainer>
        <Input
          type="text"
          value={searchTerm}
          onChange={handleSearchChange}
          onKeyDown={handleKeyDown}
          placeholder="검색"
        />
        {showResults && (
          <ResultsContainer>
            {searchResults.map((result, index) => (
              <ResultItem
                key={index}
                onClick={() => handleAddTech(result)}
                className={index === selectedResultIndex ? "selected" : ""}>
                {result}
              </ResultItem>
            ))}
          </ResultsContainer>
        )}
      </SearchContainer>
      <ButtonContainer>
        <Button onClick={() => navigator(`/profile/${user.id}`)}>취소</Button>
        <Button
          onClick={() => {
            saveTech(userTechList);
          }}>
          저장
        </Button>
      </ButtonContainer>
    </Container>
  );
};

export default TechList;

const Container = styled.div`
  padding: 20px;
  position: relative;
`;

const SearchContainer = styled.div`
  margin-top: 20px;
  position: relative;
  width: 200px;
`;

const Input = styled.input`
  padding: 8px;
  width: 200px;
  font-size: 15px;
`;

const ResultsContainer = styled.div`
  position: absolute;
  top: 100%;
  left: 0;
  width: 100%;
  max-height: 200px;
  overflow-y: auto;
  margin-top: 5px;
  border: 2px solid black;
  background-color: #fff;
  z-index: 1;
`;

const ResultItem = styled.div`
  cursor: pointer;
  margin-bottom: 5px;
  padding: 5px;

  &.selected {
    background-color: #3498db;
    color: #fff;
  }
`;

const ButtonContainer = styled.div`
  margin-top: 20px;
  display: flex;
  justify-content: center;
`;

const Button = styled.div`
  cursor: pointer;
  padding: 10px;
  background-color: #3498db;
  color: #fff;
  margin-right: 10px;
  display: inline-block;

  &:last-child {
    margin-right: 0;
  }
`;
