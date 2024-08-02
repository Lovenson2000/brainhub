import React from "react";

type TaskTitleProps = {
  title: string;
};

export default function TaskTitle({ title }: TaskTitleProps) {
  return (
    <div className="text-2xl m-1 text-slate-900">
      <h2>{title}</h2>
    </div>
  );
}
