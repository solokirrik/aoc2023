# Solution for day 10 task2

## Slow

```cpp
template <class T> bool pt_in_polygon(const T &test, const std::vector &polygon)
{
    if (polygon.size() < 3)
        return false;

    std::vector::const_iterator end = polygon.end();

    T last_pt = polygon.back();

    last_pt.x -= test.x;
    last_pt.y -= test.y;

    double sum = 0.0;

    for (
        std::vector::const_iterator iter = polygon.begin();
        iter != end;
        ++iter)
    {
        T cur_pt = *iter;
        cur_pt.x -= test.x;
        cur_pt.y -= test.y;

        double del = last_pt.x * cur_pt.y - cur_pt.x * last_pt.y;
        double xy = cur_pt.x * last_pt.x + cur_pt.y * last_pt.y;

        sum +=
            (atan((last_pt.x * last_pt.x + last_pt.y * last_pt.y - xy) / del) +
             atan((cur_pt.x * cur_pt.x + cur_pt.y * cur_pt.y - xy) / del));

        last_pt = cur_pt;
    }

    return fabs(sum) > eps;
}

T – тип точки, например : struct PointD
{
    double x, y;
};
```

## Fast

```cpp
template bool pt_in_polygon2(const T &test, const std::vector &polygon)
{

    static const int q_patt[2][2] = {{0, 1}, {3, 2}};

    if (polygon.size() < 3)
        return false;

    std::vector::const_iterator end = polygon.end();
    T pred_pt = polygon.back();
    pred_pt.x -= test.x;
    pred_pt.y -= test.y;

    int pred_q = q_patt[pred_pt.y < 0][pred_pt.x < 0];

    int w = 0;

    for (std::vector::const_iterator iter = polygon.begin(); iter != end; ++iter)
    {
        T cur_pt = *iter;

        cur_pt.x -= test.x;
        cur_pt.y -= test.y;

        int q = q_patt[cur_pt.y < 0][cur_pt.x < 0];

        switch (q - pred_q)
        {
        case -3:
            ++w;
            break;
        case 3:
            --w;
            break;
        case -2:
            if (pred_pt.x * cur_pt.y >= pred_pt.y * cur_pt.x)
                ++w;
            break;
        case 2:
            if (!(pred_pt.x * cur_pt.y >= pred_pt.y * cur_pt.x))
                --w;
            break;
        }

        pred_pt = cur_pt;
        pred_q = q;
    }

    return w != 0;
}
```
