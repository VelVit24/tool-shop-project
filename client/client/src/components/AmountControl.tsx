interface Props {
  value: number;
  min?: number;
  max?: number;
  onChange: (value: number) => void;
}
import { Minus, Plus } from 'lucide-react';
import { useState, useEffect } from 'react';

interface Props {
  value: number;
  min?: number;
  max?: number;
  onChange: (value: number) => void;
}

export default function AmountControl({
  value,
  min = 1,
  max = Infinity,
  onChange,
}: Props) {
  const [localValue, setLocalValue] = useState(value.toString());

  useEffect(() => {
    setLocalValue(value.toString());
  }, [value]);

  function clamp(val: number) {
    return Math.min(max, Math.max(min, val));
  }

  function apply(val: number) {
    const safe = clamp(val);
    setLocalValue(safe.toString());
    onChange(safe);
  }

  return (
    <div className="flex items-center gap-2">
      <button
        className="rounded-full w-10 h-10 flex items-center justify-center hover:bg-blue-300 active:bg-blue-400"
        onClick={() => apply(value - 1)}
      >
        <Minus size={20} />
      </button>

      <input
        value={localValue}
        className="
          w-12 border border-gray-200
          text-xl text-center rounded-md px-2
        "
        onChange={(e) => setLocalValue(e.target.value)}
        onBlur={() => {
          const num = Number(localValue);

          if (Number.isNaN(num)) {
            apply(min);
            return;
          }

          apply(num);
        }}
      />

      <button
        className="
        rounded-full w-10 h-10 flex items-center 
        justify-center hover:bg-blue-300 
        active:bg-blue-400"
        onClick={() => apply(value + 1)}
      >
        <Plus size={20} />
      </button>
    </div>
  );
}
