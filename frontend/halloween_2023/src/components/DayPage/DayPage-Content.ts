export default interface DayPageContent {
    DayNumb:number;
    Description: string;
}

const DayOne : DayPageContent = {
    DayNumb:1,
    Description:"Day One"
}

const DayTwo : DayPageContent = {
    DayNumb: 2,
    Description : "Day Two"
}

const DayThree: DayPageContent = {
    DayNumb: 3,
    Description: "Day Three"
}

const DayFour: DayPageContent = {
    DayNumb: 4,
    Description: "Day Four"
}

export const DayPagesContent : DayPageContent[] = [DayOne, DayTwo, DayThree, DayFour]