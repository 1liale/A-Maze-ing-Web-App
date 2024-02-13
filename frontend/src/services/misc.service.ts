export const SCORE_MULT = 10;
export const MOVE_PENALTY = 1.5;

export const computeScore = (time: number, area: number, moves: number): number => {
  console.log({ time, area, moves });
  return Math.max(Math.round((SCORE_MULT * area) / time - moves * MOVE_PENALTY), 0);
};
